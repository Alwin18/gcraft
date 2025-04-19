package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

var configTemplates string

func createTemplateFile(dst, templatePath, dstPath, moduleName string) error {
	// Parsing template file
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("gagal memparsing template: %v", err)
	}

	// Persiapkan data untuk template
	data := struct {
		ModuleName string
	}{
		ModuleName: moduleName,
	}

	// Tentukan file tujuan
	dstFile := filepath.Join(dst, dstPath)

	// ✅ Pastikan direktori tujuan tersedia
	if err := os.MkdirAll(filepath.Dir(dstFile), 0755); err != nil {
		return fmt.Errorf("gagal membuat direktori %s: %v", filepath.Dir(dstFile), err)
	}

	// Buat file
	file, err := os.Create(dstFile)
	if err != nil {
		return fmt.Errorf("gagal membuat file %s: %v", dstPath, err)
	}
	defer file.Close()

	// Eksekusi template dan tulis ke dalam file tujuan
	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("gagal mengeksekusi template: %v", err)
	}

	fmt.Printf("✅ %s berhasil dibuat di: %s\n", dstPath, dstFile)
	return nil
}

func renderTemplates(srcTemplate, dst, moduleName string) error {
	return filepath.Walk(srcTemplate, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcTemplate, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		// Buat folder
		if info.IsDir() {
			return os.MkdirAll(dstPath, 0755)
		}

		// Kalau file .tmpl, render ke file .go (atau sejenis)
		if filepath.Ext(path) == ".tmpl" {
			dstPath = dstPath[:len(dstPath)-len(".tmpl")]
			return createTemplateFile(dst, path, dstPath[len(dst)+1:], moduleName)
		}

		// Untuk file biasa (non-templating), salin langsung
		input, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			return err
		}
		return os.WriteFile(dstPath, input, 0644)
	})
}

func writeGoMod(projectPath, projectName string) error {
	goModPath := filepath.Join(projectPath, "go.mod")
	content := fmt.Sprintf("module %s\n\ngo 1.23.0\n", projectName)
	return os.WriteFile(goModPath, []byte(content), 0644)
}

var createCmd = &cobra.Command{
	Use:   "create [nama-proyek]",
	Short: "Membuat proyek baru dari template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		dst := filepath.Join(".", projectName)

		// Hapus folder jika sudah ada
		if _, err := os.Stat(dst); err == nil {
			if err := os.RemoveAll(dst); err != nil {
				fmt.Println("❌ Gagal menghapus folder lama:", err)
				return
			}
		}

		// Buat ulang folder proyek
		if err := os.MkdirAll(dst, 0755); err != nil {
			fmt.Println("❌ Gagal membuat folder:", err)
			return
		}

		// Salin template folder dasar
		srcTemplate := filepath.Join("templates", "basic-go")

		// Render semua .go.tmpl file di srcTemplate ke dst
		if err := renderTemplates(srcTemplate, dst, projectName); err != nil {
			fmt.Println("❌ Gagal merender templates:", err)
			return
		}

		// Generate go.mod sesuai nama project
		if err := writeGoMod(dst, projectName); err != nil {
			fmt.Println("❌ Gagal menulis go.mod:", err)
			return
		}

		fmt.Println("\n✅ Proyek berhasil dibuat di:", projectName)
		fmt.Println("👉 Selanjutnya:")
		fmt.Println("   cd", projectName)
		fmt.Println("   go mod tidy   # untuk mengunduh semua dependencies")
		fmt.Println("   go run main.go   # untuk menjalankan aplikasi")
	},
}

func init() {
	createCmd.Flags().StringVar(&configTemplates, "with-config", "", "Daftar file konfigurasi yang ingin disertakan, pisahkan dengan koma (misalnya fiber,gorm)")
	rootCmd.AddCommand(createCmd)
}
