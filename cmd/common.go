package cmd

import (
	"fmt"
	"github.com/nyudlts/go-aspace/lib"
	"os"
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

type UIDS struct {
	Dash   string
	Period string
}

func NormalizeUnitIds(id0 string, id1 string, id2 string, id3 string) UIDS {
	dash := id0
	period := id0

	if id1 != "" {
		dash = fmt.Sprintf("%s_%s", dash, id1)
		period = fmt.Sprintf("%s.%s", period, id1)
	}

	if id2 != "" {
		dash = fmt.Sprintf("%s_%s", dash, id2)
		period = fmt.Sprintf("%s.%s", period, id2)
	}

	if id3 != "" {
		dash = fmt.Sprintf("%s_%s", dash, id3)
		period = fmt.Sprintf("%s.%s", period, id3)
	}

	return UIDS{dash, period}

}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
