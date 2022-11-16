package aspace

import (
	"encoding/json"
	"fmt"
	"io"
)

func (a *ASClient) GetDigitalObjectIDs(repositoryId int) ([]int, error) {
	daoIds := []int{}
	endpoint := fmt.Sprintf("/repositories/%d/digital_objects?all_ids=true", repositoryId)
	response, err := a.get(endpoint, true)
	if err != nil {
		return daoIds, err
	}
	body, _ := io.ReadAll(response.Body)
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
	body, err := io.ReadAll(response.Body)
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

	body, err = io.ReadAll(response.Body)
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

	body, err = io.ReadAll(response.Body)
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
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return string(body), err
	}
	return string(body), nil
}

func (a *ASClient) GetRandomDigitalObject() (int, int, error) {
	var repositoryID = 0
	var digitalObjectID = 0
	repositoryIDs, err := a.GetRepositories()
	if err != nil {
		return repositoryID, digitalObjectID, err
	}

	repositoryID = repositoryIDs[rGen.Intn(len(repositoryIDs))]

	digitalObjectIDs, err := a.GetDigitalObjectIDs(repositoryID)
	if err != nil {
		return repositoryID, digitalObjectID, err
	}

	digitalObjectID = digitalObjectIDs[rGen.Intn(len(digitalObjectIDs))]

	return repositoryID, digitalObjectID, nil
}

func (a *ASClient) GetDigitalObjectIDsForResource(repositoryId int, resourceId int) ([]string, error) {
	doURIs := []string{}

	aoURIs, err := a.GetArchivalObjectsForResource(repositoryId, resourceId, "digital_object")

	if err != nil {
		return doURIs, err
	}

	for _, aoURI := range aoURIs {
		_, aoId, _ := URISplit(aoURI)
		ao, err := a.GetArchivalObject(repositoryId, aoId)
		if err != nil {
			return doURIs, err
		}

		for _, instance := range ao.Instances {
			if instance.InstanceType == "digital_object" {
				doURIs = append(doURIs, instance.DigitalObject["ref"])
			}
		}
	}
	return doURIs, nil
}

func (do DigitalObject) ContainsUseStatement(role string) bool {
	for _, version := range do.FileVersions {
		if version.UseStatement == role {
			return true
		}
	}
	return false
}
