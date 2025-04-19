package fs

import (
	"os"
	"text/template"
)

func RenderTemplateFile(srcPath, dstPath string, data any) error {
	tmpl, err := template.ParseFiles(srcPath)
	if err != nil {
		return err
	}

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	return tmpl.Execute(dstFile, data)
}
