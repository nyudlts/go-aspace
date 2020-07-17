package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().StringVarP(&uri, "uri", "u", "/", "endpoint uri")
	getCmd.PersistentFlags().StringVarP(&location, "location", "l", "/tmp", "location to write json")
	getCmd.PersistentFlags().StringVarP(&filename, "filename", "f", "output.json", "name of file to write")
}

var getCmd = &cobra.Command{
	Use: "get",
	Short: "Get JSON",
	Long: "Serialize JSON from ASpace api to file",
	Run: func(cmd *cobra.Command, args []string) {
		//doublecheck the location
		if dirExists(location) == false {
			fmt.Println("  ! location does not exist, defaulting to /tmp")
			location = "/tmp"
		}

		//create the output file
		outputFilePath := filepath.Join(location, filename)
		outputFile, err := os.Create(outputFilePath)
		if err != nil {
			fmt.Printf("Could not create file at: %s, exiting\n", outputFile)
			panic(err)
		}
		defer outputFile.Close()


		//request the json
		response, err := aspace.GetEndpoint(uri)
		if err != nil {
			panic(err)
		}

		body := string(response)

		w := bufio.NewWriter(outputFile)
		n, err := w.WriteString(body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d %d\n", len(response), n)
		w.Flush()

	},

}