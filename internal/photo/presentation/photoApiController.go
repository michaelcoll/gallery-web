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
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

var rangeRxp = regexp.MustCompile(`(?P<Unit>.*)=(?P<Start>[0-9]+)-(?P<End>[0-9]*)`)

func (c *ApiController) mediaList(ctx *gin.Context) {
	daemon, err := c.getDaemonById(ctx, true)
	if err != nil {
		handleError(ctx, err)
		return
	}

	start, end, err := extractRangeHeader(ctx.GetHeader("Range"))
	if err != nil {
		handleError(ctx, err)
		return
	}

	photos, total, err := c.photoService.List(ctx.Request.Context(), daemon, uint32(start), uint32(end-start))
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Header("Content-Range", fmt.Sprintf("%s %d-%d/%d", "photo", start, start+len(photos), total))
	ctx.JSON(http.StatusOK, photos)
}

func extractRangeHeader(rangeHeader string) (int, int, error) {
	r := rangeRxp.FindStringSubmatch(rangeHeader)
	st := http.StatusRequestedRangeNotSatisfiable

	if len(r) < 4 {
		return 0, 0, Errorf(st, "Range is not valid, supported format : photo=0-25")
	}

	if r[1] != "photo" {
		return 0, 0, Errorf(st, "Unit in range is not valid, supported unit : photo")
	}

	start, errStart := strconv.Atoi(r[2])
	end, errEnd := strconv.Atoi(r[3])

	if len(r[3]) == 0 {
		end = 0
	}

	if errStart != nil {
		return 0, 0, Errorf(st, "Start range is not valid")
	}

	if len(r[3]) != 0 && errEnd != nil {
		return 0, 0, Errorf(st, "End range is not valid")
	}

	if end != 0 && start >= end {
		return 0, 0, Errorf(st, "Range is not valid, start > end")
	}

	return start, end, nil
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
