package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	createCmd.Flags().StringP("module", "m", "", "Go module name (default: project name)")
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(makeCmd)
}

var rootCmd = &cobra.Command{
	Use:   "gcraft",
	Short: "gcraft - Go project scaffolding CLI",
	Long:  "gcraft adalah CLI untuk generate project Golang berdasarkan template",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}
