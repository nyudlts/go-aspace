package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	rootCmd.AddCommand(tsvCmd)
	tsvCmd.PersistentFlags().StringVarP(&input, "input", "i", "/tmp/export.tsv", "input tsv file")
	tsvCmd.MarkFlagRequired("input")
	tsvCmd.PersistentFlags().StringVarP(&location, "location", "l", "/tmp", "Location to write EAD File")
}

var tsvCmd = &cobra.Command{
	Use:   "tsv",
	Short: "export resources from archivesspace from tsv file",
	Long:  `export resources from archivesspace from tsv file`,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		exportTsv()
		elapsed := time.Since(start)
		log.Printf("  * export took %s", elapsed)
	},
}

func exportTsv() {
	if !dirExists(location) {
		fmt.Println("  * location '%s' does not exist, defaulting to /tmp")
	}

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Split(line, "\t")
		repId, err := strconv.Atoi(cols[0])
		if err != nil {
			panic(err)
		}
		resId, err := strconv.Atoi(cols[1])
		if err != nil {
			panic(err)
		}
		resource, err := aspace.GetResourceByID(repId, resId)
		if err != nil {
			panic(err)
		}

		uId := NormalizeUnitIds(resource.ID0, resource.ID1, resource.ID2, resource.ID3)
		fmt.Printf("  * Exporting %s %s\n", uId.Period, resource.Title)
		err = aspace.SerializeEAD(repId, resId, location, true, false, false, false, false, uId.Dash)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
