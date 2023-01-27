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
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var statusMapping = map[codes.Code]int{
	codes.NotFound:         http.StatusNotFound,
	codes.InvalidArgument:  http.StatusBadRequest,
	codes.Unavailable:      http.StatusServiceUnavailable,
	codes.PermissionDenied: http.StatusForbidden,
	codes.OutOfRange:       http.StatusRequestedRangeNotSatisfiable,
}

func handleError(ctx *gin.Context, err error) {
	st := getStatus(err)

	if st == http.StatusInternalServerError {
		panic(err)
	}

	ctx.JSON(st, gin.H{"message": err.Error()})
}

func getStatus(err error) int {
	st, _ := status.FromError(err)

	if httpStatus, exists := statusMapping[st.Code()]; exists {
		return httpStatus
	}

	if httpStatus, ok := FromError(err); ok {
		return httpStatus
	}

	return http.StatusInternalServerError
}
