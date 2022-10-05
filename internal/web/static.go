package web

import "embed"

//go:embed dist/index.html
//go:embed dist/favicon.ico
//go:embed dist/assets
var Static embed.FS
