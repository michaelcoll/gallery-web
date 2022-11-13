/*
 * Copyright (c) 2022 Michaël COLL.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package caller

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	photov1 "github.com/michaelcoll/gallery-proto/gen/proto/go/photo/v1"

	"github.com/michaelcoll/gallery-web/internal/photo/domain/model"
)

type PhotoServiceGrpcCaller struct {
}

func New() *PhotoServiceGrpcCaller {
	return &PhotoServiceGrpcCaller{}
}

func (c *PhotoServiceGrpcCaller) List(ctx context.Context, d *model.Daemon, page uint32, pageSize uint32) ([]*model.Photo, error) {

	client, conn, err := createClient(d)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "can't connect to the daemon (%v)", err)
	}
	defer closeConnection(conn)

	photoResponse, err := client.GetPhotos(ctx, &photov1.GetPhotosRequest{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, err
	}

	photos := make([]*model.Photo, len(photoResponse.GetPhotos()))

	for i, photo := range photoResponse.GetPhotos() {
		photos[i] = toDomain(photo)
	}

	return photos, nil
}

func (c *PhotoServiceGrpcCaller) Exists(ctx context.Context, d *model.Daemon, hash string) (bool, error) {

	client, conn, err := createClient(d)
	if err != nil {
		return false, status.Errorf(codes.Unavailable, "can't connect to the daemon (%v)", err)
	}
	defer closeConnection(conn)

	resp, err := client.ExistsByHash(ctx, &photov1.ExistsByHashRequest{
		Hash: hash,
	})
	if err != nil {
		return false, err
	}

	return resp.Exists, nil
}

func (c *PhotoServiceGrpcCaller) ContentByHash(ctx context.Context, d *model.Daemon, hash string) ([]byte, string, error) {

	client, conn, err := createClient(d)
	if err != nil {
		return nil, "", status.Errorf(codes.Unavailable, "can't connect to the daemon (%v)", err)
	}
	defer closeConnection(conn)

	stream, err := client.ContentByHash(ctx, &photov1.ContentByHashRequest{
		Hash: hash,
	})
	if err != nil {
		return nil, "", err
	}

	content := make([]byte, 0)
	var contentType string
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			return content, contentType, nil
		}
		if err != nil {
			return nil, "", err
		}

		content = append(content, chunk.Data...)
		contentType = chunk.ContentType
	}
}

func (c *PhotoServiceGrpcCaller) ThumbnailByHash(ctx context.Context, d *model.Daemon, hash string, width uint32, height uint32) ([]byte, string, error) {

	client, conn, err := createClient(d)
	if err != nil {
		return nil, "", status.Errorf(codes.Unavailable, "can't connect to the daemon (%v)", err)
	}
	defer closeConnection(conn)

	stream, err := client.ThumbnailByHash(ctx, &photov1.ThumbnailByHashRequest{
		Hash:   hash,
		Width:  width,
		Height: height,
	})
	if err != nil {
		return nil, "", err
	}

	content := make([]byte, 0)
	var contentType string
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			return content, contentType, nil
		}
		if err != nil {
			return nil, "", err
		}

		content = append(content, chunk.Data...)
		contentType = chunk.ContentType
	}
}

func createClient(d *model.Daemon) (photov1.PhotoServiceClient, *grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	daemonAddr := fmt.Sprintf("%s:%d", d.Hostname, d.Port)

	conn, err := grpc.Dial(daemonAddr, opts...)
	if err != nil {
		return nil, nil, err
	}
	client := photov1.NewPhotoServiceClient(conn)

	return client, conn, nil
}

func closeConnection(conn *grpc.ClientConn) {
	err := conn.Close()
	if err != nil {
		log.Fatalf("fail to close the daemon connection : %v", err)
	}
}
