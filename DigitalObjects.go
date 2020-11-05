package aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (a *ASClient) GetDigitalObjectsByRepositoryId(repositoryId int) ([]int, error) {
	daoIds := []int{}
	endpoint := fmt.Sprintf("/repositories/%d/digital_objects?all_ids=true", repositoryId)
	response, err := a.get(endpoint, true)
	if err != nil {
		return daoIds, err
	}
	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &daoIds)
	if err != nil {
		return daoIds, err
	}

	return daoIds, nil
}

func (a *ASClient) GetDigitalObject(repositoryId int, daoId int) (DigitalObject, error) {
	do := DigitalObject{}
	endpoint := fmt.Sprintf("/repositories/%d/digital_objects/%d", repositoryId, daoId)
	response, err := a.get(endpoint, true)
	if err != nil {
		return do, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return do, err
	}
	err = json.Unmarshal(body, &do)
	if err != nil {
		return do, err
	}

	return do, nil
}

func (a *ASClient) UpdateDigitalObject(repositoryId int, daoId int, dao DigitalObject) (string, error) {
	endpoint := fmt.Sprintf("/repositories/%d/digital_objects/%d", repositoryId, daoId)
	body, err := json.Marshal(dao)
	if err != nil {
		return "", err
	}
	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return "", err
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (a *ASClient) CreateDigitalObject(repositoryId int, dao DigitalObject) (string, error) {
	endpoint := fmt.Sprintf("/repositories/%d/digital_objects", repositoryId)
	body, err := json.Marshal(dao)
	if err != nil {
		return "", err
	}
	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return "", err
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (a *ASClient) DeleteDigitalObject(repositoryId int, doId int) (string, error) {
	endpoint := fmt.Sprintf("/repositories/%d/digital_objects/%d", repositoryId, doId)
	response, err := a.delete(endpoint)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return string(body), err
	}
	return string(body), nil
}
