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

package service

import "github.com/michaelcoll/gallery-web/internal/photo/domain/model"

type PhotoService struct {
	c DaemonCaller
}

func New(c DaemonCaller) PhotoService {
	return PhotoService{c: c}
}

func (s *PhotoService) List() ([]*model.Photo, error) {
	return s.c.List()
}

func (s *PhotoService) GetByHash(hash string) (*model.Photo, error) {
	return s.c.GetByHash(hash)
}

func (s *PhotoService) ContentByHash(hash string) ([]byte, error) {
	return s.c.ContentByHash(hash)
}
