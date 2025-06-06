package templates

import (
	"embed"
	"io/fs"
)

//go:embed all:basic-go
var BasicGoTemplate embed.FS

//go:embed all:handler
var HandlerTemplate embed.FS

//go:embed all:service
var ServiceTemplate embed.FS

// GetBasicGoTemplate returns the embedded basic-go template filesystem
func GetBasicGoTemplate() fs.FS {
	subFS, err := fs.Sub(BasicGoTemplate, "basic-go")
	if err != nil {
		panic(err)
	}
	return subFS
}

// GetHandlerTemplate returns the embedded handler template filesystem
func GetHandlerTemplate() fs.FS {
	subFS, err := fs.Sub(HandlerTemplate, "handler")
	if err != nil {
		panic(err)
	}
	return subFS
}

// GetServiceTemplate returns the embedded handler template filesystem
func GetServiceTemplate() fs.FS {
	subFS, err := fs.Sub(ServiceTemplate, "service")
	if err != nil {
		panic(err)
	}
	return subFS
}