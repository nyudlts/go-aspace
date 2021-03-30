package aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (a *ASClient) GetTopContainerIDs(repositoryID int) ([]int, error) {
	var topContainers = []int{}
	endpoint := fmt.Sprintf("/repositories/%d/top_containers?all_ids=true", repositoryID)
	response, err := a.get(endpoint, true)
	if err != nil {
		return topContainers, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return topContainers, err
	}
	err = json.Unmarshal(body, &topContainers)
	if err != nil {
		return topContainers, err
	}

	return topContainers, nil
}
