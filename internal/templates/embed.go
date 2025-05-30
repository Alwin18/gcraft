package templates

import (
	"embed"
	"io/fs"
)

//go:embed all:basic-go
var BasicGoTemplate embed.FS

// GetBasicGoTemplate returns the embedded basic-go template filesystem
func GetBasicGoTemplate() fs.FS {
	subFS, err := fs.Sub(BasicGoTemplate, "basic-go")
	if err != nil {
		panic(err)
	}
	return subFS
}
