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
)

func (c *ApiController) mediaList(ctx *gin.Context) {
	daemon, err := c.getDaemonById(ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}

	photos, err := c.photoService.List(ctx.Request.Context(), daemon)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, photos)
}

func (c *ApiController) getByHash(ctx *gin.Context) {
	daemon, err := c.getDaemonById(ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}

	hash := ctx.Param("hash")
	photo, err := c.photoService.GetByHash(ctx.Request.Context(), daemon, hash)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

func (c *ApiController) contentByHash(ctx *gin.Context) {
	daemon, err := c.getDaemonById(ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}

	hash := ctx.Param("hash")
	photoContent, contentType, err := c.photoService.ContentByHash(ctx.Request.Context(), daemon, hash)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Data(http.StatusOK, contentType, photoContent)
}
