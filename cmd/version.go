package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "v1.0.3"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of gcraft",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gcraft version %s\n", Version)
	},
}
