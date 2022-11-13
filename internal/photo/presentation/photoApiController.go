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
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *ApiController) mediaList(ctx *gin.Context) {
	daemon, err := c.getDaemonById(ctx, true)
	if err != nil {
		handleError(ctx, err)
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "0"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "25"))

	photos, err := c.photoService.List(ctx.Request.Context(), daemon, uint32(page), uint32(pageSize))
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, photos)
}

func (c *ApiController) contentByHash(ctx *gin.Context) {

	daemon, err := c.getDaemonById(ctx, false)
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

func (c *ApiController) thumbnailByHash(ctx *gin.Context) {

	daemon, err := c.getDaemonById(ctx, false)
	if err != nil {
		handleError(ctx, err)
		return
	}

	hash := ctx.Param("hash")
	w, _ := strconv.Atoi(ctx.DefaultQuery("width", "0"))
	h, _ := strconv.Atoi(ctx.DefaultQuery("height", "0"))
	photoContent, contentType, err := c.photoService.ThumbnailByHash(ctx.Request.Context(), daemon, hash, uint32(w), uint32(h))
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Data(http.StatusOK, contentType, photoContent)
}
