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

func writeGoMod(projectPath, projectName string) error {
	goModPath := filepath.Join(projectPath, "go.mod")
	content := fmt.Sprintf("module %s\n\ngo 1.23.0\n", projectName)
	return os.WriteFile(goModPath, []byte(content), 0644)
}

func GetModuleName() (string, error) {
	goModPath := "go.mod"

	// Baca file go.mod
	content, err := os.ReadFile(goModPath)
	if err != nil {
		return "", fmt.Errorf("failed to read go.mod: %w", err)
	}

	// Cari baris yang dimulai dengan "module"
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module ") {
			// Ambil nama modul setelah "module"
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	return "", fmt.Errorf("module name not found in go.mod")
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
		return processFile(templateFS, path, targetPath, data, projectName)
	})
}

func processFile(templateFS fs.FS, srcPath, targetPath string, data TemplateData, projectName string) error {
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

	if err := writeGoMod(filepath.Join(".", projectName), projectName); err != nil {
		fmt.Println("❌ Gagal menulis go.mod:", err)
		return err
	}

	fmt.Printf("Created: %s\n", targetPath)
	return nil
}

func CreateHandlerStructure(name string) error {
	templateFS := templates.GetHandlerTemplate()

	return fs.WalkDir(templateFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip root directory
		if path == "." {
			return nil
		}

		if d.IsDir() {
			// Skip directories, as we only want to process files
			return nil
		}

		moduleName, err := GetModuleName()
		if err != nil {
			return fmt.Errorf("failed to get module name: %w", err)
		}

		// Define the target file path
		targetFile := filepath.Join("internal", "handlers", fmt.Sprintf("%s.go", name))

		// Process the template file
		return ProcessTemplateFile(templateFS, path, targetFile, TemplateData{
			ProjectName: name,
			ModuleName:  moduleName,
		})
	})
}

func CreateServiceStructure(name string) error {
	templateFS := templates.GetServiceTemplate()

	return fs.WalkDir(templateFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip root directory
		if path == "." {
			return nil
		}

		if d.IsDir() {
			// Skip directories, as we only want to process files
			return nil
		}

		moduleName, err := GetModuleName()
		if err != nil {
			return fmt.Errorf("failed to get module name: %w", err)
		}

		// Define the target file path
		targetFile := filepath.Join("internal", "services", fmt.Sprint(name), fmt.Sprintf("%s.go", name))

		// Process the template file
		return ProcessTemplateFile(templateFS, path, targetFile, TemplateData{
			ProjectName: name,
			ModuleName:  moduleName,
		})
	})
}

func ProcessTemplateFile(fsys fs.FS, templatePath string, targetFile string, data TemplateData) error {
	// Convert target file path to lowercase
	targetFile = strings.ToLower(targetFile)

	fmt.Printf("Reading template: %s\n", templatePath)
	tmplBytes, err := fs.ReadFile(fsys, templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template file %s: %w", templatePath, err)
	}

	tmpl, err := template.New(filepath.Base(templatePath)).Parse(string(tmplBytes))
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", templatePath, err)
	}

	fmt.Printf("Creating directory: %s\n", filepath.Dir(targetFile))
	if err := os.MkdirAll(filepath.Dir(targetFile), 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", filepath.Dir(targetFile), err)
	}

	fmt.Printf("Creating file: %s\n", targetFile)
	outFile, err := os.Create(targetFile)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", targetFile, err)
	}
	defer outFile.Close()

	// Capitalize the first letter of each word in the template data fields
	data.ProjectName = CapitalizeFirst(data.ProjectName)

	fmt.Printf("Executing template: %s -> %s\n", templatePath, targetFile)
	if err := tmpl.Execute(outFile, data); err != nil {
		return fmt.Errorf("failed to execute template %s: %w", templatePath, err)
	}

	return nil
}

func CapitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
