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
	photov1 "github.com/michaelcoll/gallery-proto/gen/proto/go/photo/v1"
	"github.com/michaelcoll/gallery-web/internal/photo/domain/model"
)

func toDomain(grpc *photov1.Photo) *model.Photo {
	return &model.Photo{
		Hash: grpc.Hash,
		Path: grpc.Path,

		DateTime:     grpc.DateTime,
		Iso:          grpc.Iso,
		ExposureTime: grpc.ExposureTime,
		XDimension:   grpc.XDimension,
		YDimension:   grpc.YDimension,
		Model:        grpc.Model,
		FNumber:      grpc.FNumber,
	}
}
