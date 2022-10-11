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
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/michaelcoll/gallery-web/internal/photo/domain/service"
)

const apiPort = ":8080"

type PhotoController struct {
	s service.PhotoService
}

func NewPhotoController(s service.PhotoService) PhotoController {
	return PhotoController{s: s}
}

func (c *PhotoController) Serve() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	serveStatic(router)

	router.GET("/api/media", c.list)
	router.GET("/api/media/:hash", c.getByHash)
	router.GET("/api/media/:hash/content", c.contentByHash)

	// Listen and serve on 0.0.0.0:8080
	fmt.Printf("%s Listening API on 0.0.0.0%s\n", color.GreenString("✓"), color.GreenString(apiPort))
	err := router.Run(apiPort)
	if err != nil {
		log.Fatalf("Error starting server : %v", err)
	}
}

func (c *PhotoController) list(ctx *gin.Context) {
	photos, err := c.s.List(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't obtain the media list !"})
	}

	ctx.JSON(http.StatusOK, photos)
}

func (c *PhotoController) getByHash(ctx *gin.Context) {
	hash := ctx.Param("hash")

	photo, err := c.s.GetByHash(ctx.Request.Context(), hash)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't obtain the media !"})
	}

	ctx.JSON(http.StatusOK, photo)
}

func (c *PhotoController) contentByHash(ctx *gin.Context) {
	hash := ctx.Param("hash")

	photoContent, err := c.s.ContentByHash(ctx.Request.Context(), hash)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't obtain the media !"})
	}

	ctx.Data(http.StatusOK, "image/jpeg", photoContent)
}
