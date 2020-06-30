package cmd

import (
	"github.com/nyudlts/go-aspace/lib"
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(reportCmd)
	reportCmd.PersistentFlags().StringVarP(&repositories, "repositories", "r", "2", "List of repository ids to be included in sample set")

}
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "report",
	Long:  `report`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("  **	exporting report")
		report()
	},
}

func report() {
	outputTSV, err := os.Create("report.tsv")
	if err != nil {
		panic(err)
	}
	defer outputTSV.Close()
	writer := bufio.NewWriter(outputTSV)
	writer.WriteString("repository\tresource\ttitle\turi\tpublished\tcreators\n")
	fmt.Println(" ** Getting Resources")
	repositoryIds := splitRepos(repositories)
	for _, repositoryId := range repositoryIds {
		fmt.Println("  ** Processing repository", repositoryId)
		resourceIds, err := aspace.GetResourceIDsByRepository(repositoryId)
		if err != nil {
			panic(err)
		}
		for _, resourceId := range resourceIds {
			fmt.Println("  ** Processing resource ", resourceId, " in repository ", repositoryId)
			resource, err := aspace.GetResourceByID(repositoryId, resourceId)
			if err != nil {
				panic(err)
			}
			creators, err := getCreators(resource.LinkedAgents)
			if err != nil {
				panic(err)
			}
			entry := fmt.Sprintf("%d\t%d\t%s\t%s\t%v\t%v\n", repositoryId, resourceId, resource.Title, resource.URI, resource.Publish, creators)
			writer.WriteString(entry)

		}
	}

	fmt.Println("Complete")

}

func getCreators(agents []*lib.LinkedAgent) ([]string, error) {
	creators := []string{}
	for _, agent := range agents {
		if agent.Role == "creator" {
			creators = append(creators, agent.Ref)
		}

	}
	return creators, nil
}
