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

import (
	"context"

	"github.com/michaelcoll/gallery-web/internal/photo/domain/model"
)

type PhotoServiceCaller interface {
	List(ctx context.Context, d *model.Daemon, page uint32, pageSize uint32) ([]*model.Photo, error)
	Exists(ctx context.Context, d *model.Daemon, hash string) (bool, error)
	ContentByHash(ctx context.Context, d *model.Daemon, hash string) ([]byte, string, error)
	ThumbnailByHash(ctx context.Context, d *model.Daemon, hash string, width uint32, height uint32) ([]byte, string, error)
}
