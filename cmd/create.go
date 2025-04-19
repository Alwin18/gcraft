package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Alwin18/gracft/internal/fs"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [nama-proyek]",
	Short: "Membuat proyek baru",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		dst := filepath.Join(".", projectName)

		// Jika folder sudah ada, hapus dulu
		if _, err := os.Stat(dst); err == nil {
			fmt.Println("📁 Folder sudah ada, menghapus folder lama:", dst)
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

		// Salin template
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

		fmt.Println("✅ Proyek berhasil dibuat:", projectName)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func writeGoMod(projectPath, projectName string) error {
	goModPath := filepath.Join(projectPath, "go.mod")
	content := fmt.Sprintf("module %s\n\ngo 1.20\n", "github.com/"+projectName)
	return os.WriteFile(goModPath, []byte(content), 0644)
}
