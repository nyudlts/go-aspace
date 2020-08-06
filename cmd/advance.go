package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(adQuery)
}

var adQuery = &cobra.Command{
	Use: "adquery",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("HI")
	},
}
