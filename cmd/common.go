package cmd

import (
	"fmt"
	"github.com/nyudlts/go-aspace/lib"
	"strconv"
	"strings"
)

var aspace = lib.Client

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
