package aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

func (a *ASClient) GetRepositoryList() ([]int, error) {
	repIds := []int{}
	endpoint := "/repositories"
	response, err := a.get(endpoint, false)
	if err != nil {
		return repIds, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return repIds, err
	}

	reps := make([]map[string]interface{}, 1, 1)
	err = json.Unmarshal(body, &reps)

	for i := range reps {
		rep := fmt.Sprintf("%v", reps[i]["uri"])
		repId, err := strconv.Atoi(rep[len(rep)-1:])
		if err != nil {
			return repIds, err
		}
		repIds = append(repIds, repId)
	}

	return repIds, nil
}
