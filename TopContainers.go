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

func (a *ASClient) GetTopContainer(repositoryID int, topContainerID int) (TopContainer, error) {
	tc := TopContainer{}
	endpoint := fmt.Sprintf("/repositories/%d/top_containers/%d", repositoryID, topContainerID)

	reponse, err := a.get(endpoint, true)
	if err != nil {
		return tc, err
	}

	body, err := ioutil.ReadAll(reponse.Body)
	if err != nil {
		return tc, err
	}

	err = json.Unmarshal(body, &tc)
	if err != nil {
		return tc, err
	}

	return tc, nil
}

func (a *ASClient) GetTopContainerIDsForResource(repositoryID int, resourceID int) ([]string, error) {
	tcs := []string{}
	responseMap := []map[string]string{}
	endpoint := fmt.Sprintf("/repositories/%d/resources/%d/top_containers", repositoryID, resourceID)

	reponse, err := a.get(endpoint, true)
	if err != nil {
		return tcs, err
	}

	body, err := ioutil.ReadAll(reponse.Body)
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

func (a *ASClient) GetTopContainersForResource(repositoryID int, resourceID int) (map[string]TopContainer, error) {
	tcs := map[string]TopContainer{}
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
		tcs[tc.Indicator] = tc
	}
	return tcs, nil
}
