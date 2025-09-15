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
	doURI := fmt.Sprintf("/repositories/%d/digital_objects/%d", repositoryId, daoId)

	return a.GetDigitalObjectFromURI(doURI)
}

func (a *ASClient) GetDigitalObjectFromURI(doURI string) (DigitalObject, error) {
	do := DigitalObject{}
	response, err := a.get(doURI, true)
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

func (a *ASClient) UpdateDigitalObject(repositoryId int, daoId int, dao *DigitalObject) (*APIResponse, error) {
	endpoint := fmt.Sprintf("/repositories/%d/digital_objects/%d", repositoryId, daoId)
	body, err := json.Marshal(dao)
	if err != nil {
		return nil, err
	}
	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return nil, err
	}

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	apiResponse := &APIResponse{}
	if err := json.Unmarshal(body, apiResponse); err != nil {
		return nil, err
	}

	return apiResponse, nil
}

func (a *ASClient) CreateDigitalObject(repositoryId int, dao *DigitalObject) (*APIResponse, error) {
	endpoint := fmt.Sprintf("/repositories/%d/digital_objects", repositoryId)
	body, err := json.Marshal(dao)
	if err != nil {
		return nil, err
	}
	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return nil, err
	}

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	apiResponse := &APIResponse{}
	if err := json.Unmarshal(body, apiResponse); err != nil {
		return nil, err
	}

	return apiResponse, nil
}

func (a *ASClient) DeleteDigitalObject(repositoryId int, doId int) (*APIResponse, error) {
	doURI := fmt.Sprintf("/repositories/%d/digital_objects/%d", repositoryId, doId)

	return a.DeleteDigitalObjectFromURI(doURI)
}

func (a *ASClient) DeleteDigitalObjectFromURI(doURI string) (*APIResponse, error) {
	response, err := a.delete(doURI)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	apiResponse := &APIResponse{}
	if err := json.Unmarshal(body, apiResponse); err != nil {
		return nil, err
	}
	return apiResponse, nil
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

	aoURIs, err := a.GetArchivalObjectsForResource(repositoryId, resourceId)

	if err != nil {
		return doURIs, err
	}

	for _, aoURI := range aoURIs {
		aspaceURI, err := ParseAspaceURI(aoURI)
		if err != nil {
			return doURIs, err
		}

		ao, err := a.GetArchivalObject(repositoryId, aspaceURI.ObjectID)
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

func (a *ASClient) GetDigitalObjectIDsForArchivalObjectFromURI(aoURI string) ([]string, error) {
	doURIs := []string{}

	ao, err := a.GetArchivalObjectFromURI(aoURI)
	if err != nil {
		return doURIs, err
	}

	for _, instance := range ao.Instances {
		if instance.InstanceType == "digital_object" {
			doURIs = append(doURIs, instance.DigitalObject["ref"])
		}
	}

	return doURIs, nil
}
