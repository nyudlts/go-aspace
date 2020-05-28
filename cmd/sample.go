package cmd

import (

	"fmt"
	"github.com/nyudlts/go-aspace/lib"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var size int
var xportlocation string
var repositories string

func init() {
	client = lib.Client
	rootCmd.AddCommand(sampleCmd)

	exportCmd.Flags().IntVarP(&size, "size", "s", 0, "Size of the sample")
	exportCmd.Flags().StringVarP(&repositories, "repositories", "z", "2", "List of repository ids to be included in sample set")
	exportCmd.Flags().StringVarP(&xportlocation, "xportlocation", "x", "/tmp", "Location to write EAD Files")

}

var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "sample resources from archivesspace",
	Long:  `sample resources from archivesspace`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("* generating sample set")

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
		ids, err := client.GetResourceIDsByRepository(r)
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
		ead := fmt.Sprintf("%s/%d_%d.xml", location, r.Rep, r.Res)
		fmt.Printf("* serializeing %s\n", ead)
		err := client.SerializeEAD(r.Rep, r.Res, location, true, false, false, false, false)
		if err != nil {
			fmt.Printf("** ERROR: Could not serialize %s ✗\n", ead)
			fmt.Printf(err.Error())
			break
		}
		fmt.Println(" ✓")
	}

}

func dirExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func splitRepos(s string) []int {
	repos := []int{}
	a := strings.Split(s, " ")
	for _, i := range a {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(fmt.Errorf("%s is not a valid repository id", s))
		}
		repos = append(repos, j)
	}
	return repos
}