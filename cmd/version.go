package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)
const VERSION = "0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "prints current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version %s\n", VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
