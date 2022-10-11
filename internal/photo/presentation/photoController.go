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
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/michaelcoll/gallery-web/internal/photo/domain/model"
	"github.com/michaelcoll/gallery-web/internal/photo/domain/service"
)

const apiPort = ":8080"

type PhotoController struct {
	photoService  *service.PhotoService
	daemonService *service.DaemonService
}

func NewPhotoController(s *service.PhotoService, d *service.DaemonService) PhotoController {
	return PhotoController{photoService: s, daemonService: d}
}

func (c *PhotoController) Serve() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	serveStatic(router)

	router.GET("/api/daemon", c.daemonList)
	router.GET("/api/daemon/:id", c.daemonById)
	router.GET("/api/daemon/:id/media", c.mediaList)
	router.GET("/api/daemon/:id/media/:hash", c.getByHash)
	router.GET("/api/daemon/:id/media/:hash/content", c.contentByHash)

	// Listen and serve on 0.0.0.0:8080
	fmt.Printf("%s Listening API on 0.0.0.0%s\n", color.GreenString("✓"), color.GreenString(apiPort))
	err := router.Run(apiPort)
	if err != nil {
		log.Fatalf("Error starting server : %v", err)
	}
}

func (c *PhotoController) daemonList(ctx *gin.Context) {
	list := c.daemonService.List()

	ctx.JSON(http.StatusOK, list)
}

func (c *PhotoController) daemonById(ctx *gin.Context) {
	daemon, err := c.getDaemonById(ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, daemon)
}

func (c *PhotoController) mediaList(ctx *gin.Context) {
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

func (c *PhotoController) getByHash(ctx *gin.Context) {
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

func (c *PhotoController) contentByHash(ctx *gin.Context) {
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

func (c *PhotoController) getDaemonById(ctx *gin.Context) (*model.Daemon, error) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "daemon id not valid (%s)", id)
	}
	daemon, err := c.daemonService.ById(id)
	if err != nil {
		return nil, err
	}

	return daemon, nil
}
