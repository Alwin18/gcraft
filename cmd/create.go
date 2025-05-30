package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Alwin18/gcraft/internal/fs"
	"github.com/spf13/cobra"
)

var Version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of gcraft",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gcraft version %s\n", Version)
	},
}

var createCmd = &cobra.Command{
	Use:   "create [project-name]",
	Short: "Create a new Go project from template",
	Long:  "Create a new Go project structure based on the basic-go template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		moduleName, _ := cmd.Flags().GetString("module")

		// If module name not provided, use project name
		if moduleName == "" {
			moduleName = projectName
		}

		// Check if directory already exists
		if _, err := os.Stat(projectName); !os.IsNotExist(err) {
			fmt.Printf("Error: Directory '%s' already exists\n", projectName)
			os.Exit(1)
		}

		// Create project
		fmt.Printf("Creating project '%s' with module '%s'...\n", projectName, moduleName)

		if err := fs.ProcessTemplate(projectName, moduleName); err != nil {
			fmt.Printf("Error creating project: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\n✅ Project '%s' created successfully!\n", projectName)
		fmt.Printf("📂 Location: %s\n", filepath.Join(".", projectName))
		fmt.Printf("\nNext steps:\n")
		fmt.Printf("  cd %s\n", projectName)
		fmt.Printf("  go mod tidy\n")
		fmt.Printf("  go run cmd/main.go\n")
	},
}

func init() {
	createCmd.Flags().StringP("module", "m", "", "Go module name (default: project name)")
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(versionCmd)
}
