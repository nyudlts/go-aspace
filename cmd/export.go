package cmd

import (
	"fmt"
	"github.com/nyudlts/go-aspace/lib"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var resourceId int
var repositoryId int
var location string
var client lib.ASClient

func init() {
	client = lib.Client
	rootCmd.AddCommand(exportCmd)
	exportCmd.PersistentFlags().IntVar(&repositoryId, "repositoryId", 0, "Id of the repository")
	exportCmd.PersistentFlags().IntVar(&resourceId, "resourceId", 0, "Id of the resource")
	exportCmd.PersistentFlags().StringVar(&location, "location", "/tmp", "Location to write EAD File")
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export a resource from archivesspace",
	Long:  `export a resource from archivesspace`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("* exporting %s/%d_%d.xml\n", location, repositoryId, resourceId)
		err := exportEAD(); if err != nil {
			panic(err)
		}
		//check file exists
		if !checkExists() {
			panic("Export failed")
		}

		fmt.Println("* export complete")
	},
}

func exportEAD() error {
	//ensure the location exists
	if _, err := os.Stat(location); os.IsNotExist(err) {
		return fmt.Errorf("%v does not exist, exiting", err.Error())
	}

	//ensure the repository exists
	repos, err := client.GetRepositoryList()
	if err != nil {
	  return err
	}

	if !contains(repos, repositoryId) {
		return fmt.Errorf("Repository ID %d does not exist in the current ArchivesSpace instance", repositoryId)
	}

	//serialize the EAD
	err = client.SerializeEAD(repositoryId, resourceId, location); if err != nil {
		return err
	}

	return nil
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func checkExists() bool {
	fn := fmt.Sprintf("%d_%d.xml", repositoryId, resourceId)
	f := filepath.Join(location, fn)
	if _, err := os.Stat(f); os.IsNotExist(err) {
		return false
	}
	return true
}

