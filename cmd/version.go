package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "get version of go-aspace",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-aspace v0.1")
	},
}
