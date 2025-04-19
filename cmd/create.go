package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Alwin18/gracft/internal/fs"
	"github.com/spf13/cobra"
)

var configTemplates string

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
		if err := fs.CopyDir(srcTemplate, dst); err != nil {
			fmt.Println("❌ Gagal menyalin template:", err)
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
