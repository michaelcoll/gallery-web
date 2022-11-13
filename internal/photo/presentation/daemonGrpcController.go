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

package presentation

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"google.golang.org/grpc"

	daemonv1 "github.com/michaelcoll/gallery-proto/gen/proto/go/daemon/v1"

	"github.com/michaelcoll/gallery-web/internal/photo/domain/model"
	"github.com/michaelcoll/gallery-web/internal/photo/domain/service"
)

const (
	grpcPort = ":9000"
)

type GrpcController struct {
	s *service.DaemonService

	daemonv1.UnimplementedDaemonServiceServer
}

func NewGrpcController(s *service.DaemonService) GrpcController {
	return GrpcController{s: s}
}

func (c *GrpcController) Serve() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	grpcServer := grpc.NewServer()
	daemonv1.RegisterDaemonServiceServer(grpcServer, c)

	go c.s.Watch()

	fmt.Printf("%s Listening daemons on 0.0.0.0%s\n", color.GreenString("✓"), color.GreenString(grpcPort))
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (c *GrpcController) Register(_ context.Context, req *daemonv1.RegisterRequest) (*daemonv1.RegisterResponse, error) {

	id, expiresIn, err := c.s.Register(fromRegister(req))
	if err != nil {
		return nil, err
	}

	return &daemonv1.RegisterResponse{
		Uuid:  id.String(),
		ExpIn: expiresIn,
	}, nil
}

func (c *GrpcController) HeartBeat(_ context.Context, req *daemonv1.HeartBeatRequest) (*daemonv1.HeartBeatResponse, error) {
	daemon, err := fromHeartBeat(req)
	if err != nil {
		return nil, err
	}

	err = c.s.HeartBeat(daemon)
	if err != nil {
		return nil, err
	}

	return &daemonv1.HeartBeatResponse{}, nil
}

func fromRegister(request *daemonv1.RegisterRequest) *model.Daemon {
	return &model.Daemon{
		Id:       uuid.New(),
		Name:     request.DaemonName,
		Owner:    request.DaemonOwner,
		Hostname: request.DaemonHost,
		Port:     int(request.DaemonPort),
		Version:  request.DaemonVersion,
		Alive:    true,
		New:      true,
		LastSeen: time.Now(),
	}
}

func fromHeartBeat(request *daemonv1.HeartBeatRequest) (*model.Daemon, error) {
	id, err := uuid.Parse(request.Uuid)
	if err != nil {
		return nil, err
	}

	return &model.Daemon{
		Id:       id,
		Name:     request.DaemonName,
		Owner:    request.DaemonOwner,
		Hostname: request.DaemonHost,
		Port:     int(request.DaemonPort),
		Version:  request.DaemonVersion,
		Alive:    false,
		LastSeen: time.Now(),
	}, nil
}
