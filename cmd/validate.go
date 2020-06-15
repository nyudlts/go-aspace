package cmd

import (
	"fmt"
	"github.com/nyudlts/go-aspace/lib"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
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
		fmt.Printf("  * Validating repositories: %s\n", repositories)
		validateResources()
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
		resourceIds, err := aspace.GetResourceIDsByRepository(repositoryId)
		if err != nil {
			panic(err)
		}


		for _, resourceId := range resourceIds {

			resource, err := aspace.GetResourceByID(repositoryId, resourceId)
			if err != nil {
				panic(err)
			}

			if (published == true && resource.Publish == true) || published == false {
				err := validateEAD(repositoryId, resourceId)

				if err != nil {
					logFile.WriteString(fmt.Sprintf("%d\t%d\t%s\n", repositoryId, resourceId, resource.Title))
					fmt.Printf("  ✗ Validation failed for %d %d %s  ✗\n", repositoryId, resourceId, resource.Title)
				} else {
					fmt.Printf("  ✓ Validation was succesful for %d %d %s  ✓\n", repositoryId, resourceId, resource.Title)
				}

			} else if published == true && resource.Publish == false {
				fmt.Printf("  ! Skipping validation for %d %d %s  !\n", repositoryId, resourceId, resource.Title)
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
