package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

)

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.PersistentFlags().IntVarP(&repositoryId, "repositoryId", "r", 0, "Id of the repository")
	searchCmd.PersistentFlags().StringVarP(&query, "query", "q", ".", "Query String")
	searchCmd.PersistentFlags().StringVarP(&searchType, "type", "t", "resource", "Type of search [resource, accession]")
}

var hits = []string{}
var totalHits int = 0



var searchCmd = &cobra.Command{
	Use: "search",

	Run: func(cmd *cobra.Command, args []string) {
		BasicSearch(1)
		for i,hit := range hits {
			fmt.Println(i, ".", hit)
		}
	},

}

func BasicSearch(page int) {

	results, err := aspace.Search(repositoryId, searchType, query, page)

	if err != nil {
		panic(err)
	}

	if page == 1 { totalHits = results.TotalHits }

	for _, hit := range results.Results {

		hits = append(hits, fmt.Sprintf("\t%v\t%v\t%v", hit["identifier"], hit["title"], hit["uri"]))
	}

	if results.ThisPage < results.LastPage {
		BasicSearch(page + 1)
	}

}

