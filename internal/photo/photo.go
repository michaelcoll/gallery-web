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

package photo

import (
	"github.com/michaelcoll/gallery-web/internal/photo/domain/service"
	"github.com/michaelcoll/gallery-web/internal/photo/infrastructure/caller"
	"github.com/michaelcoll/gallery-web/internal/photo/presentation"
)

type Module struct {
	s          service.PhotoService
	photoCtrl  presentation.ApiController
	daemonCtrl presentation.GrpcController
}

func (m *Module) GetPhotoController() *presentation.ApiController {
	return &m.photoCtrl
}

func (m *Module) GetDaemonController() *presentation.GrpcController {
	return &m.daemonCtrl
}

func (m *Module) GetService() *service.PhotoService {
	return &m.s
}

func New() Module {
	photoServiceCaller := caller.New()
	photoService := service.NewPhotoService(photoServiceCaller)
	daemonService := service.NewDaemonService(photoServiceCaller)

	return Module{
		s:          photoService,
		photoCtrl:  presentation.NewApiController(&photoService, &daemonService),
		daemonCtrl: presentation.NewGrpcController(&daemonService),
	}
}
