package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use: "go-aspace",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from go-aspace")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

