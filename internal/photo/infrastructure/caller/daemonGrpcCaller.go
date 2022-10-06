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
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	photov1 "github.com/michaelcoll/gallery-proto/gen/proto/go/photo/v1"
	"github.com/michaelcoll/gallery-web/internal/photo/domain/model"
)

const (
	daemonHost = "localhost"
	daemonPort = 9000
)

type DaemonGrpcCaller struct {
}

func New() *DaemonGrpcCaller {
	return &DaemonGrpcCaller{}
}

func (c *DaemonGrpcCaller) List() ([]*model.Photo, error) {

	client, conn := createClient(daemonHost, daemonPort)
	defer closeConnection(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	photoResponse, err := client.GetPhotos(ctx, &photov1.GetPhotosRequest{})
	if err != nil {
		return nil, err
	}

	photos := make([]*model.Photo, len(photoResponse.GetPhotos()))

	for i, photo := range photoResponse.GetPhotos() {
		photos[i] = toDomain(photo)
	}

	return photos, nil
}

func (c *DaemonGrpcCaller) GetByHash(hash string) (*model.Photo, error) {

	client, conn := createClient(daemonHost, daemonPort)
	defer closeConnection(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetByHash(ctx, &photov1.GetByHashRequest{
		Hash: hash,
	})
	if err != nil {
		return nil, err
	}

	return toDomain(resp.Photo), nil
}

func (c *DaemonGrpcCaller) ContentByHash(hash string) ([]byte, error) {
	client, conn := createClient(daemonHost, daemonPort)
	defer closeConnection(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.ContentByHash(ctx, &photov1.ContentByHashRequest{
		Hash: hash,
	})
	if err != nil {
		return nil, err
	}

	content := make([]byte, 0)
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			return content, nil
		}
		if err != nil {
			log.Fatalf("cannot receive %v", err)
		}

		content = append(content, chunk.Data...)
	}

	return nil, nil
}

func createClient(daemonHost string, daemonPort int) (photov1.PhotoServiceClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	daemonAddr := fmt.Sprintf("%s:%d", daemonHost, daemonPort)

	conn, err := grpc.Dial(daemonAddr, opts...)
	if err != nil {
		log.Fatalf("fail to contact the daemon : %v", err)
	}
	client := photov1.NewPhotoServiceClient(conn)

	return client, conn
}

func closeConnection(conn *grpc.ClientConn) {
	err := conn.Close()
	if err != nil {
		log.Fatalf("fail to close the daemon connection : %v", err)
	}
}
