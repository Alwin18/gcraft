package fs

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Alwin18/gcraft/internal/templates"
)

type TemplateData struct {
	ProjectName string
	ModuleName  string
	// Tambahkan field lain sesuai kebutuhan
}

// ProcessTemplate processes template files and creates project structure
func ProcessTemplate(projectName, moduleName string) error {
	templateFS := templates.GetBasicGoTemplate()

	data := TemplateData{
		ProjectName: projectName,
		ModuleName:  moduleName,
	}

	return fs.WalkDir(templateFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip root directory
		if path == "." {
			return nil
		}

		// Create target path
		targetPath := filepath.Join(projectName, path)

		if d.IsDir() {
			// Create directory
			return os.MkdirAll(targetPath, 0755)
		}

		// Process file
		return processFile(templateFS, path, targetPath, data)
	})
}

func processFile(templateFS fs.FS, srcPath, targetPath string, data TemplateData) error {
	// Read template file
	content, err := fs.ReadFile(templateFS, srcPath)
	if err != nil {
		return fmt.Errorf("failed to read template file %s: %w", srcPath, err)
	}

	// Remove .tmpl extension from target path
	if strings.HasSuffix(targetPath, ".tmpl") {
		targetPath = strings.TrimSuffix(targetPath, ".tmpl")
	}

	// Ensure target directory exists
	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// If file has .tmpl extension, process as template
	if strings.HasSuffix(srcPath, ".tmpl") {
		tmpl, err := template.New(filepath.Base(srcPath)).Parse(string(content))
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", srcPath, err)
		}

		// Create target file
		file, err := os.Create(targetPath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", targetPath, err)
		}
		defer file.Close()

		// Execute template
		if err := tmpl.Execute(file, data); err != nil {
			return fmt.Errorf("failed to execute template %s: %w", srcPath, err)
		}
	} else {
		// Copy file as-is
		if err := os.WriteFile(targetPath, content, 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", targetPath, err)
		}
	}

	fmt.Printf("Created: %s\n", targetPath)
	return nil
}
