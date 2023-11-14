package aspace

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

// Return a slice of Repository IDs for the instance
func (a *ASClient) GetRepositories() ([]int, error) {
	repIds := []int{}
	endpoint := "/repositories"
	response, err := a.get(endpoint, false)
	if err != nil {
		return repIds, err
	}
	body, err := io.ReadAll(response.Body)
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

// Return a Repository struct for given Repository ID
func (a *ASClient) GetRepository(repositoryID int) (Repository, error) {
	repository := Repository{}
	endpoint := fmt.Sprintf("/repositories/%d", repositoryID)
	response, err := a.get(endpoint, true)
	if err != nil {
		return repository, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return repository, err
	}

	err = json.Unmarshal(body, &repository)
	if err != nil {
		return repository, err
	}

	return repository, nil
}

// Get a random Repository ID
func (a *ASClient) GetRandomRepository() (int, error) {

	repositoryIDs := []int{2, 3, 6}

	return repositoryIDs[rGen.Intn(len(repositoryIDs))], nil
}
