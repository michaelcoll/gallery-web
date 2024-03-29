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
	"time"

	cachecontrol "github.com/joeig/gin-cachecontrol"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/michaelcoll/gallery-web/internal/photo/domain/service"
)

const apiPort = ":8080"

type ApiController struct {
	photoService  *service.PhotoService
	daemonService *service.DaemonService
}

func NewApiController(s *service.PhotoService, d *service.DaemonService) ApiController {
	return ApiController{photoService: s, daemonService: d}
}

func (c *ApiController) Serve() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	serveStatic(router)
	addCommonMiddlewares(router)

	public := router.Group("/api/v1")
	mediaGroup := router.Group("/api/v1")
	private := router.Group("/api/v1")

	addJWTMiddlewares(private)
	mediaGroup.Use(cachecontrol.New(&cachecontrol.Config{
		Public:    true,
		MaxAge:    cachecontrol.Duration(7 * 24 * time.Hour),
		Immutable: true,
	}))

	private.GET("/daemon", c.daemonList)
	private.GET("/daemon/:id", c.daemonById)
	private.GET("/daemon/:id/media", c.mediaList)

	public.GET("/daemon/:id/status", c.daemonStatusById)

	mediaGroup.GET("/daemon/:id/media/:hash", c.contentByHash)
	mediaGroup.GET("/daemon/:id/thumbnail/:hash", c.thumbnailByHash)

	// Listen and serve on 0.0.0.0:8080
	fmt.Printf("%s Listening API on http://0.0.0.0%s\n", color.GreenString("✓"), color.GreenString(apiPort))
	err := router.Run(apiPort)
	if err != nil {
		log.Fatalf("Error starting server : %v", err)
	}
}
