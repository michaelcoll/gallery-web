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

package model

type Photo struct {
	Hash string `json:"hash"`
	Path string `json:"path"`

	DateTime     string `json:"dateTime,omitempty"`
	Iso          int32  `json:"iso,omitempty"`
	ExposureTime string `json:"exposureTime,omitempty"`
	XDimension   int32  `json:"xDimension,omitempty"`
	YDimension   int32  `json:"yDimension,omitempty"`
	Model        string `json:"model,omitempty"`
	FNumber      string `json:"fNumber,omitempty"`
}
