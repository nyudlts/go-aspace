package aspace

import (
	"encoding/json"
	"fmt"
	"io"
)

func (a *ASClient) GetTopContainerIDs(repositoryID int) ([]int, error) {
	var topContainers = []int{}
	endpoint := fmt.Sprintf("/repositories/%d/top_containers?all_ids=true", repositoryID)
	response, err := a.get(endpoint, true)
	if err != nil {
		return topContainers, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return topContainers, err
	}
	err = json.Unmarshal(body, &topContainers)
	if err != nil {
		return topContainers, err
	}

	return topContainers, nil
}

func (a *ASClient) GetTopContainer(repositoryID int, topContainerID int) (TopContainer, error) {
	tc := TopContainer{}
	endpoint := fmt.Sprintf("/repositories/%d/top_containers/%d", repositoryID, topContainerID)

	reponse, err := a.get(endpoint, true)
	if err != nil {
		return tc, err
	}

	body, err := io.ReadAll(reponse.Body)
	if err != nil {
		return tc, err
	}

	err = json.Unmarshal(body, &tc)
	if err != nil {
		return tc, err
	}

	return tc, nil
}

// Update a Top Container for a given Repository and Accession ID
func (a *ASClient) UpdateTopContainer(repositoryID int, topContainerID int, topContainer TopContainer) (string, error) {
	endpoint := fmt.Sprintf("/repositories/%d/top_containers/%d", repositoryID, topContainerID)
	body, err := json.Marshal(topContainer)
	if err != nil {
		return "", err
	}
	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return "", err
	}

	msg, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(msg), nil
}

// Delete a Top Container
func (a *ASClient) DeleteTopContainer(repositoryID int, topContainerID int) (string, error) {
	endpoint := fmt.Sprintf("/repositories/%d/top_containers/%d", repositoryID, topContainerID)
	response, err := a.delete(endpoint)
	if err != nil {
		return fmt.Sprintf("code %d", response.StatusCode), err
	}
	msg, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Sprintf("code %d", response.StatusCode), err
	}

	return string(msg), nil
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

func (a *ASClient) GetTopContainersForResource(repositoryID int, resourceID int) ([]TopContainer, error) {
	tcs := []TopContainer{}
	tcIds, err := a.GetTopContainerIDsForResource(repositoryID, resourceID)
	if err != nil {
		return tcs, err
	}
	for _, tcId := range tcIds {
		_, tcId, err := URISplit(tcId)
		if err != nil {
			return tcs, err
		}

		tc, err := a.GetTopContainer(repositoryID, tcId)
		if err != nil {
			return tcs, err
		}
		tcs = append(tcs, tc)
	}
	return tcs, nil
}
