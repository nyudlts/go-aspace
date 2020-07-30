package cmd

import (
	"fmt"
	"github.com/nyudlts/go-aspace/lib"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.PersistentFlags().IntVarP(&repositoryId, "repositoryId", "r", 2, "Id of the repository")
	searchCmd.PersistentFlags().StringVarP(&query, "query", "q", ".", "Query String")
	searchCmd.PersistentFlags().StringVarP(&searchType, "type", "t", "", "Type of search [resource, accession]")
	searchCmd.PersistentFlags().StringVarP(&fieldList, "fieldList", "f", "", "Field List")
	searchCmd.PersistentFlags().IntVarP(&pageLimit, "pageLimit", "p", 1, "page limit of search")
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

	q := lib.QueryString{Query: ""}
	q.AddParameter( "q", query)
	q.AddParameter("page", strconv.Itoa(page))
	if searchType != "" { q.AddParameter("type[]", searchType) }
	if fieldList != "" { q.AddParameter("fl",fieldList ) }
	var fields []string
	if fieldList == "" {
		fields = []string{"uri"}
	} else {
		fields = strings.Split(fieldList, ",")
	}

	results, err := aspace.Search(repositoryId, q.Query)

	if err != nil {
		panic(err)
	}

	if page == 1 { totalHits = results.TotalHits }

	for _, hit := range results.Results {

		h := ""
		for _, field := range fields {
			if h == "" {
				h = fmt.Sprintf("%s", hit[field])
			} else {
				h = fmt.Sprintf("%s\t%s", h, hit[field])
			}
		}
		hits = append(hits, h)
	}

	if(results.ThisPage < pageLimit) {
		if(results.ThisPage < results.LastPage) {
			BasicSearch(page + 1)
		}
	}


}

