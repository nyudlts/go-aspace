package cmd

import (
	"fmt"
	"github.com/nyudlts/go-aspace"
	"github.com/spf13/cobra"
	"io/ioutil"
	"math/rand"
	"time"
)

func init() {
	rootCmd.AddCommand(sampleCmd)
	sampleCmd.PersistentFlags().IntVarP(&size, "size", "s", 1, "Size of the sample set")
	sampleCmd.PersistentFlags().StringVarP(&repositories, "repositories", "r", "2", "List of repository ids to be included in sample set")
	sampleCmd.PersistentFlags().StringVarP(&location, "location", "l", "/tmp", "Location to write EAD Files")
}

var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "Generate an EAD sample set",
	Long:  `Generate a sample set of resources as EAD from ArchivesSpace`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("  * generating sample set")
		sample()
	},
}

type RepRes struct {
	Rep int
	Res int
}

func sample() {

	if !dirExists(location) {
		fmt.Println("* location '%s' does not exist, defaulting to /tmp")
	}

	repositoryIds := splitRepos(repositories)

	repres := []RepRes{}

	for _, r := range repositoryIds {
		ids, err := aspace.GetResourceIDsByRepository(r)
		if err != nil {
			fmt.Println("ERROR: Could not retreive list of resources for repository id %d, skipping\n", r)
		}

		for _, i := range ids {
			repres = append(repres, RepRes{r, i})
		}
	}

	represSize := len(repres)

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= size; i++ {
		r := repres[rand.Intn(represSize)]
		filename := fmt.Sprintf("%d_%d", r.Rep, r.Res)
		eadPath := fmt.Sprintf("%s/%s.xml", location, filename)
		fmt.Printf("  * serializing %s\n", eadPath)
		err := aspace.SerializeEAD(r.Rep, r.Res, location, true, false, false, false, false, filename)
		if err != nil {
			fmt.Println("  ✗ ERROR: Could not serialize %s ✗\n", eadPath)
			fmt.Printf(err.Error())
			break
		}
		ead, err := ioutil.ReadFile(eadPath)
		if err != nil {
			fmt.Printf("  ✗ Validator could not open ead file: %v✗\n", err)
		}
		err = main.ValidateEAD(ead)
		if err != nil {
			fmt.Printf("  ✗ validation Failed, check output file in an XML editor ✗\n%v\n", err)
		} else {
			fmt.Printf("  ✓ %s is valid ead, export complete ✓\n", eadPath)
		}
	}

}
