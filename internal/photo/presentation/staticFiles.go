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
	"github.com/michaelcoll/gallery-web/internal/web"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func serveStatic(router *gin.Engine) {
	static, _ := fs.Sub(web.Static, "dist")
	staticAssets, _ := fs.Sub(web.Static, "dist/assets")
	staticImg, _ := fs.Sub(web.Static, "dist/img")
	staticIndexFS(http.FS(static), router)
	router.StaticFS("/assets", http.FS(staticAssets))
	router.StaticFS("/img", http.FS(staticImg))
	router.StaticFS("/apple-touch-icon.png", http.FS(static))
	router.StaticFS("/icon-192.png", http.FS(static))
	router.StaticFS("/icon-512.png", http.FS(static))
	router.StaticFS("/favicon.ico", http.FS(static))
	router.StaticFS("/manifest.webmanifest", http.FS(static))
}

func staticIndexFS(fs http.FileSystem, router *gin.Engine) {
	relativePath := "/"
	handler := createStaticHandler(relativePath, fs)

	// Register GET and HEAD handlers
	router.GET(relativePath, handler)
	router.HEAD(relativePath, handler)
}

func createStaticHandler(relativePath string, fs http.FileSystem) gin.HandlerFunc {
	fileServer := http.StripPrefix(relativePath, http.FileServer(fs))

	return func(c *gin.Context) {
		// Check if file exists and/or if we have permission to access it
		f, err := fs.Open("/")
		if err != nil {
			c.Writer.WriteHeader(http.StatusNotFound)
			return
		}
		f.Close()

		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}
