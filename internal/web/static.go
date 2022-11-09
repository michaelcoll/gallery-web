package web

import "embed"

//go:embed dist/assets
//go:embed dist/img
//go:embed dist/apple-touch-icon.png
//go:embed dist/favicon.ico
//go:embed dist/icon-192.png
//go:embed dist/icon-512.png
//go:embed dist/index.html
//go:embed dist/manifest.webmanifest
var Static embed.FS
