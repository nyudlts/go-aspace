package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type Query struct {
	Query FieldQuery `json:"query"`
}

type FieldQuery struct {
	Field 		string 	`json:"field"`
	Value 		string 	`json:"value"`
	JsonModelType 	string `json:"jsonmodel_type"`
	Negated 	bool 	`json:"negated"`
	Literal 	bool 	`json:"literal"`
}

func init() {
	rootCmd.AddCommand(adQueryCmd)
	adQueryCmd.PersistentFlags().IntVarP(&repositoryId, "repository-id", "r", 6, "resource id")
	adQueryCmd.PersistentFlags().StringVarP(&queryField, "query-field", "f", "*",  "query field")
	adQueryCmd.PersistentFlags().StringVarP(&value, "value", "v", "*", "query value")
	adQueryCmd.PersistentFlags().StringVarP(&queryType, "query-type", "q", "resource", "query type")
}

var hits []string = []string{}
var tsvWriter *bufio.Writer

var adQueryCmd = &cobra.Command{
	Use: "adquery",
	Run: func(cmd *cobra.Command, args []string) {

		fq := FieldQuery{
			Field:         queryField,
			Value:         value,
			JsonModelType: "field_query",
			Negated:       false,
			Literal:       true,
		}
		query := Query{Query:fq}

		jsonQuery, err := json.Marshal(query)
		if err != nil {
			panic(query)
		}

		adQuery := string(jsonQuery)
		pageLimit = 100

		tsv, err := os.Create("output.tsv")
		if err != nil {
			panic(err)
		}
		defer tsv.Close()
		tsvWriter  = bufio.NewWriter(tsv)
		tsvWriter.WriteString("model_type\turi\ttitle\n")
		advancedQuery(adQuery, 1)
		tsvWriter.Flush()

	},

}

func advancedQuery(query string, page int) {

	results, err := aspace.AdvancedSearch(page, repositoryId, queryType, query)

	if err != nil {
		panic(err)
	}

	if page == 1 {
		fmt.Println("# of matches:", results.TotalHits)
	}
	for _, hit := range results.Results {
		tsvWriter.WriteString(fmt.Sprintf("%s\t%s\t%s\n", hit["jsonmodel_type"], hit["uri"],hit["title"]))
		tsvWriter.Flush()
	}

	if !(results.ThisPage >= pageLimit) {
		if results.ThisPage < results.LastPage {
			advancedQuery(query, page + 1)
		}
	}


}
