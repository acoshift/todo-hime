package template

import (
	"embed"
)

//go:embed *.tmpl
//go:embed **/*.tmpl
var FS embed.FS
