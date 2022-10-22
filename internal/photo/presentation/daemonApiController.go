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
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/michaelcoll/gallery-web/internal/photo/domain/model"
)

func (c *ApiController) daemonList(ctx *gin.Context) {
	list := c.daemonService.List(extractEmailFromToken(ctx))

	ctx.JSON(http.StatusOK, list)
}

func (c *ApiController) daemonById(ctx *gin.Context) {
	daemon, err := c.getDaemonById(ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, daemon)
}

func (c *ApiController) getDaemonById(ctx *gin.Context) (*model.Daemon, error) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "daemon id not valid (%s)", id)
	}
	daemon, err := c.daemonService.ById(id)
	if err != nil {
		return nil, err
	}
	if daemon.Owner != extractEmailFromToken(ctx) {
		return nil, status.Error(codes.NotFound, "daemon not found")
	}

	return daemon, nil
}
