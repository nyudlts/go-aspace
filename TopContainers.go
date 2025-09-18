package aspace

import (
	"encoding/json"
	"fmt"
	"io"
)

func (a *ASClient) CreateTopContainer(repositoryID int, topContainer *TopContainer) (*APIResponse, error) {
	endpoint := fmt.Sprintf("/repositories/%d/top_containers", repositoryID)
	body, err := json.Marshal(topContainer)
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
	err = json.Unmarshal(responseBody, apiResponse)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal API response: %v", err)
	}

	return apiResponse, nil
}

func (a *ASClient) GetTopContainerIDs(repositoryID int) ([]int, error) {
	var topContainers = []int{}
	endpoint := fmt.Sprintf("/repositories/%d/top_containers?all_ids=true", repositoryID)
	response, err := a.get(endpoint, true)
	if err != nil {
		return topContainers, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return topContainers, err
	}
	if err = json.Unmarshal(body, &topContainers); err != nil {
		return topContainers, err
	}

	return topContainers, nil
}

func (a *ASClient) GetTopContainer(repositoryID int, topContainerID int) (*TopContainer, error) {

	endpoint := fmt.Sprintf("/repositories/%d/top_containers/%d", repositoryID, topContainerID)

	response, err := a.get(endpoint, true)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	topContainer := &TopContainer{}
	if err := json.Unmarshal(body, topContainer); err != nil {
		return nil, err
	}

	return topContainer, nil
}

// Update a Top Container for a given Repository and Accession ID
func (a *ASClient) UpdateTopContainer(repositoryID int, topContainerID int, topContainer *TopContainer) (*APIResponse, error) {

	endpoint := fmt.Sprintf("/repositories/%d/top_containers/%d", repositoryID, topContainerID)
	body, err := json.Marshal(topContainer)
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
		return nil, fmt.Errorf("could not unmarshal API response: %v", err)
	}

	return apiResponse, nil
}

// Delete a Top Container
func (a *ASClient) DeleteTopContainer(repositoryID int, topContainerID int) (*APIResponse, error) {
	endpoint := fmt.Sprintf("/repositories/%d/top_containers/%d", repositoryID, topContainerID)
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
		return nil, fmt.Errorf("could not unmarshal API response: %v", err)
	}

	return apiResponse, nil
}

func (a *ASClient) GetTopContainerIDsForResource(repositoryID int, resourceID int) ([]string, error) {
	tcs := []string{}
	responseMap := []map[string]string{}
	endpoint := fmt.Sprintf("/repositories/%d/resources/%d/top_containers", repositoryID, resourceID)

	reponse, err := a.get(endpoint, true)
	if err != nil {
		return tcs, err
	}

	body, err := io.ReadAll(reponse.Body)
	if err != nil {
		return tcs, err
	}

	err = json.Unmarshal(body, &responseMap)
	if err != nil {
		return tcs, err
	}

	for _, tc := range responseMap {
		tcs = append(tcs, tc["ref"])
	}

	return tcs, nil
}

func (a *ASClient) GetTopContainersForResource(repositoryID int, resourceID int) ([]*TopContainer, error) {
	tcs := []*TopContainer{}
	tcIds, err := a.GetTopContainerIDsForResource(repositoryID, resourceID)
	if err != nil {
		return tcs, err
	}
	for _, tcId := range tcIds {
		aspaceURI, err := ParseAspaceURI(tcId)
		if err != nil {
			return tcs, err
		}

		tc, err := a.GetTopContainer(repositoryID, aspaceURI.ObjectID)
		if err != nil {
			return tcs, err
		}
		tcs = append(tcs, tc)
	}
	return tcs, nil
}
