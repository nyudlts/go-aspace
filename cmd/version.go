package cmd

import (
	"fmt"
	"github.com/nyudlts/go-aspace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "get version of go-aspace",
	Run: func(cmd *cobra.Command, args []string) {
		version()
	},
}

func version() {
	fmt.Println("go-aspace", main.LibraryVersion)
}
