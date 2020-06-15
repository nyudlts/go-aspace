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

func NormalizeUnitIds(id0 string, id1 string, id2 string, id3 string) string {

	id := id0
	if id1 != "" {
		id = fmt.Sprintf("%s.%s", id, id1)
	}

	if id2 != "" {
		id = fmt.Sprintf("%s.%s", id, id2)
	}

	if id3 != "" {
		id = fmt.Sprintf("%s.%s", id, id3)
	}

	return id

}
