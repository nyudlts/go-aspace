package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(aoCmd)
	aoCmd.PersistentFlags().IntVarP(&repositoryId, "repositoryId", "r", 0, "Id of the repository")
	aoCmd.PersistentFlags().IntVarP(&aoId, "aoId", "a", 0, "Id of the Archival Object")
}

var aoCmd = &cobra.Command{
	Use:   "ao",
	Short: "export an ao from archivesspace",
	Long:  `export a ao from archivesspace`,
	Run: func(cmd *cobra.Command, args []string) {
		ao, err := aspace.GetArchivalObjectById(repositoryId, aoId)
		if err != nil {
			panic(err)
		}
		fmt.Println(ao.RefID)
	},
}
