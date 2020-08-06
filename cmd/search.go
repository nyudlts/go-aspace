package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.PersistentFlags().IntVarP(&repositoryId, "repositoryId", "r", 2, "Id of the repository")
	searchCmd.PersistentFlags().StringVarP(&query, "query", "q", ".", "Query String")
	searchCmd.PersistentFlags().StringVarP(&searchType, "type", "t", "resource", "Type of search [resource, accession]")
	searchCmd.PersistentFlags().IntVarP(&pageLimit, "limit", "l", 10000, "page limit")
}



type SearchResult struct {
	Identifiers interface{}
	Title		string
	URI 		string
}

func (s SearchResult) String() string {
	return fmt.Sprintf("%v\t%v\t%v", s.Identifiers, s.Title, s.URI)
}

var searchResults = []SearchResult{}
var totalHits int = 0

var searchCmd = &cobra.Command{
	Use: "search",

	Run: func(cmd *cobra.Command, args []string) {
		BasicSearch(1)
		fmt.Println("Total Hits: ", totalHits)
		for _, result := range searchResults {
			fmt.Println(result.String())
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
		sr := SearchResult {
			Identifiers: hit["identifier"],
			Title:       hit["title"].(string),
			URI:         hit["uri"].(string),
		}

		searchResults = append(searchResults, sr)
	}

	if !(results.ThisPage >= pageLimit) {
		if results.ThisPage < results.LastPage {
			BasicSearch(page + 1)
		}
	}

}

