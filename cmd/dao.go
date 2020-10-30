package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init(){
	rootCmd.AddCommand(daoCmd)
}

var daoCmd = &cobra.Command{
	Use: "dao",
	Run: func(cmd *cobra.Command, args [] string) {
		repositoryId := 2
		daoId := 4
		dao, err := aspace.GetDigitalObject(repositoryId, daoId)
		if err != nil {
			panic(err)
		}

		dao.FileVersions[0].Publish = true

		fmt.Println(dao.FileVersions[0].Identifier)


		result, err := aspace.PostDigitalObject(repositoryId, daoId, dao)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
		dao, err = aspace.GetDigitalObject(repositoryId, daoId)
		if err != nil {
			panic(err)
		}



		fmt.Println(dao.FileVersions[0].Identifier)
	},
}


