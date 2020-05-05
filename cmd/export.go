package cmd

import (
	"fmt"
	"github.com/nyudlts/go-aspace/lib"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
)

var resourceId int
var repositoryId int
var location string
var daos bool
var unpub bool
var num_cs bool
var ead3 bool
var pdf bool
var client lib.ASClient
var fn string
var validate bool
var path string

func init() {
	client = lib.Client
	rootCmd.AddCommand(exportCmd)
	exportCmd.PersistentFlags().IntVarP(&repositoryId, "repositoryId", "r", 0, "Id of the repository")
	exportCmd.MarkFlagRequired("repositoryId")
	exportCmd.PersistentFlags().IntVarP(&resourceId, "resourceId", "c", 0, "Id of the resource (collection)")
	exportCmd.MarkFlagRequired("resourceId")
	exportCmd.PersistentFlags().StringVarP(&location, "location", "l","/tmp", "Location to write EAD File")
	exportCmd.MarkFlagRequired("location")
	exportCmd.PersistentFlags().BoolVarP(&daos, "daos", "d",true, "include daos")
	exportCmd.PersistentFlags().BoolVarP(&unpub, "unpub", "u",false, "include unpublished (default false)")
	exportCmd.PersistentFlags().BoolVarP(&num_cs, "num_cs", "n",false, "include numbered components (default false)")
	exportCmd.PersistentFlags().BoolVarP(&ead3, "ead3", "e",false, "ead3 format (default false)")
	exportCmd.PersistentFlags().BoolVarP(&pdf, "pdf", "p",false, "pdf format (default false)")
	exportCmd.PersistentFlags().BoolVarP(&validate, "validate", "v",true, "validate xml (default false)")
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export a resource from archivesspace",
	Long:  `export a resource from archivesspace`,
	Run: func(cmd *cobra.Command, args []string) {

		if pdf {
			fn = fmt.Sprintf( "%d_%d.pdf", repositoryId, resourceId)
		} else {
			fn = fmt.Sprintf( "%d_%d.xml", repositoryId, resourceId)
		}

		path = filepath.Join(location, fn)

		fmt.Printf("* exporting %s/%s\n", location, fn)

		err := exportEAD(); if err != nil {
			panic(err)
		}

		//check file exists
		if !checkExists() {
			fmt.Printf("** Export failed -- exported file does not exist")
			os.Exit(1)
		}

		//validate xml
		if validate && !pdf && !ead3 {
			fmt.Println("* validating ead")
			ead, _:= ioutil.ReadFile(path)
			err := lib.ValidateEAD(ead); if err != nil {
				fmt.Printf("* validation Failed, check output file in an XML editor: %v\n", err)
			} else {
				fmt.Printf("* %s is valid ead\n", path)
			}
		}

		//exit the program
		fmt.Println("* export complete")
		os.Exit(0)
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
	err = client.SerializeEAD(repositoryId, resourceId, location, daos, unpub, num_cs, ead3, pdf); if err != nil {
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
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

