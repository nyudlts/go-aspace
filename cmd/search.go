package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search",
	Long:  `search`,
	Run:   func(cmd *cobra.Command, args []string){
		title := "Ronald L. Brown Collection on Early Instruction at the New York University School of Law"
		query := fmt.Sprintf(`{"query":{"field":"title","value":"%s", "jsonmodel_type":"field_query","negated":false,"literal":true}}`, title)
		aspace.Search(6, "resource",query)
	},
}