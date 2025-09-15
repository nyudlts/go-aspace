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
	if err != nil {
		return repIds, err
	}

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

func (a *ASClient) CreateRepository(repository *Repository) (*APIResponse, error) {
	endpoint := "/repositories"
	body, err := json.Marshal(repository)
	if err != nil {
		return nil, err
	}

	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	apiResponse := &APIResponse{}
	if err = json.Unmarshal(responseBody, apiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal API response: %v", err)
	}

	return apiResponse, nil
}

func (a *ASClient) DeleteRepository(repositoryID int) (*APIResponse, error) {
	endpoint := fmt.Sprintf("/repositories/%d", repositoryID)
	response, err := a.delete(endpoint)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	apiResponse := &APIResponse{}
	if err = json.Unmarshal(responseBody, apiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal API response: %v", err)
	}
	return apiResponse, nil
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

func (a *ASClient) UpdateRepository(repositoryID int, repository *Repository) (*APIResponse, error) {
	endpoint := fmt.Sprintf("/repositories/%d", repositoryID)
	body, err := json.Marshal(repository)
	if err != nil {
		return nil, err
	}
	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	apiResponse := &APIResponse{}
	if err = json.Unmarshal(responseBody, apiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal API response: %v", err)
	}

	return apiResponse, nil
}

// Get a random Repository ID
func (a *ASClient) GetRandomRepository() (int, error) {

	repositoryIDs, err := a.GetRepositories()
	if err != nil {
		return 0, err
	}

	return repositoryIDs[rGen.Intn(len(repositoryIDs))], nil
}
