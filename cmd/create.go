package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Alwin18/gracft/internal/fs"
	"github.com/spf13/cobra"
)

var configTemplates string

func copyConfigFiles(srcConfigDir, dstConfigDir, configTemplates string) error {
	// Ambil daftar template yang dipilih, pisahkan berdasarkan koma
	var templates []string
	if strings.TrimSpace(configTemplates) == "" {
		fmt.Println("📁 Memilih semua file konfigurasi")
		templates = []string{"app", "config", "fiber", "gorm", "loglrus", "validator"}
	} else {
		rawTemplates := strings.Split(configTemplates, ",")
		for _, t := range rawTemplates {
			if trimmed := strings.TrimSpace(t); trimmed != "" {
				templates = append(templates, trimmed)
			}
		}
	}

	// Buat folder config di dalam project jika belum ada
	if err := os.MkdirAll(dstConfigDir, 0755); err != nil {
		return fmt.Errorf("gagal membuat folder config: %v", err)
	}

	// Salin file konfigurasi yang dipilih
	for _, template := range templates {
		template = strings.TrimSpace(template)

		// Set path file sumber dan tujuan
		srcFile := filepath.Join(srcConfigDir, template+".go")
		dstFile := filepath.Join(dstConfigDir, template+".go")

		// Cek apakah file template ada
		if _, err := os.Stat(srcFile); err == nil {
			// Salin file ke folder proyek
			if err := fs.CopyDir(srcFile, dstFile); err != nil {
				return fmt.Errorf("gagal menyalin file %s: %v", template, err)
			}
		} else {
			fmt.Printf("⚠️ File konfigurasi %s tidak ditemukan.\n", srcFile)
		}
	}

	return nil
}

func writeGoMod(projectPath, projectName string) error {
	goModPath := filepath.Join(projectPath, "go.mod")
	content := fmt.Sprintf("module %s\n\ngo 1.20\n", projectName)
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
			fmt.Println("📁 Folder sudah ada, menghapus:", dst)
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
		if err := fs.CopyDir(srcTemplate, dst); err != nil {
			fmt.Println("❌ Gagal menyalin template:", err)
			return
		}

		// Salin file konfigurasi yang dipilih
		srcConfigDir := filepath.Join("templates", "snippets", "config")
		dstConfigDir := filepath.Join(dst, "config")
		if err := copyConfigFiles(srcConfigDir, dstConfigDir, configTemplates); err != nil {
			fmt.Println("❌ Gagal menyalin file konfigurasi:", err)
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
	// Flag untuk memilih file konfigurasi yang ingin disertakan
	createCmd.Flags().StringVar(&configTemplates, "with-config", "", "Daftar file konfigurasi yang ingin disertakan, pisahkan dengan koma (misalnya fiber,gorm)")
	rootCmd.AddCommand(createCmd)
}
