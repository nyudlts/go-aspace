package cmd

import (
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use: "eadid-update",
}

func Execute() error {
	return rootCmd.Execute()
}

