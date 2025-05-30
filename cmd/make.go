package cmd

import (
	"github.com/Alwin18/gcraft/internal/fs"
	"github.com/spf13/cobra"
)

var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "Generate project structure",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		fs.CreateHandlerStructure(projectName)
		fs.CreateServiceStructure(projectName)
	},
}
