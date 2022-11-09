package web

import "embed"

//go:embed dist/assets
//go:embed dist/favicon
//go:embed dist/img
//go:embed dist/index.html
var Static embed.FS
