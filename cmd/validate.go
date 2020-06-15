package cmd

import (
	"fmt"
	"github.com/nyudlts/go-aspace/lib"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"time"
)

func init() {
	client = lib.Client
	rootCmd.AddCommand(validateCmd)
	validateCmd.PersistentFlags().StringVarP(&repositories, "repositories", "r", "2", "List of repository ids to be included in sample set")
	validateCmd.PersistentFlags().StringVarP(&location, "location", "l", "/tmp", "Location to write EAD Files")
	validateCmd.PersistentFlags().StringVarP(&filename, "filename", "f", "go-aspace-validator.tsv", "Name of output file")
	validateCmd.PersistentFlags().BoolVarP(&published, "published", "p", true, "Validate only published resources")
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate all resources in a repository",
	Long:  `Validate all resources in a repository`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("  * Validating resources")
		start := time.Now()
		validateResources()
		elapsed := time.Since(start)
		log.Printf("  * validation took %s", elapsed)
	},
}

func validateResources() {

	if !dirExists(location) {
		fmt.Printf("  ! location '%s' does not exist, defaulting to /tmp\n", location)
	}

	logFile, err := createLogFile()
	defer logFile.Close()

	if err != nil {
		panic(err)
	}
	repositoryIds := splitRepos(repositories)

	for _, repositoryId := range repositoryIds {
		fmt.Printf("  * Validating resources in repository: %d\n", repositoryId)
		resourceIds, err := aspace.GetResourceIDsByRepository(repositoryId)
		if err != nil {
			panic(err)
		}

		for _, resourceId := range resourceIds {

			resource, err := aspace.GetResourceByID(repositoryId, resourceId)

			if err != nil {
				panic(err)
			}

			unitIds := NormalizeUnitIds(resource.ID0, resource.ID1, resource.ID2, resource.ID3)

			if (published == true && resource.Publish == true) || published == false {
				err := validateEAD(repositoryId, resourceId)

				if err != nil {
					logFile.WriteString(fmt.Sprintf("%d\t%d\t%s\t%s\n", repositoryId, resourceId, unitIds, resource.Title))
					fmt.Printf("  ✗ Validation failed for %s %s  ✗\n", unitIds, resource.Title)
				} else {
					fmt.Printf("  ✓ Validation was succesful for %s %s  ✓\n", unitIds, resource.Title)
				}

			} else if published == true && resource.Publish == false {
				fmt.Printf("  ! Skipping validation for %s %s  !\n", unitIds, resource.Title)
			}

		}
	}

}

func validateEAD(repId int, resId int) error {
	ead, err := lib.Client.GetEADAsByteArray(repId, resId)
	if err != nil {
		panic(err)
	}

	return lib.ValidateEAD(ead)
}

func createLogFile() (*os.File, error) {
	path := filepath.Join(location, filename)
	file, err := os.Create(path)
	return file, err
}
