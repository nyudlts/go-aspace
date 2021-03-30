package aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (a *ASClient) GetResourceIDs(repositoryId int) ([]int, error) {
	var repositoryIds []int
	endpoint := fmt.Sprintf("/repositories/%d/resources?all_ids=true", repositoryId)
	response, err := a.get(endpoint, true)
	if err != nil {
		return repositoryIds, err
	}
	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &repositoryIds)
	if err != nil {
		return repositoryIds, err
	}
	return repositoryIds, nil
}

func (a *ASClient) GetResource(repositoryId int, resourceId int) (Resource, error) {

	r := Resource{}

	endpoint := fmt.Sprintf("/repositories/%d/resources/%d", repositoryId, resourceId)
	response, err := a.get(endpoint, true)

	if err != nil {
		return r, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

func (a *ASClient) UpdateResource(repositoryId int, resourceId int, resource Resource) (string, error) {
	responseMessage := ""
	endpoint := fmt.Sprintf("/repositories/%d/resources/%d", repositoryId, resourceId)
	body, err := json.Marshal(resource)
	if err != nil {
		return responseMessage, err
	}
	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return responseMessage, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return responseMessage, err
	}

	responseMessage = string(responseBody)
	return responseMessage, nil
}

func (a *ASClient) CreateResource(repositoryId int, resource Resource) (string, error) {
	responseMessage := ""
	endpoint := fmt.Sprintf("/repositories/%d", repositoryId)
	body, err := json.Marshal(resource)
	if err != nil {
		return responseMessage, err
	}
	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return responseMessage, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return responseMessage, err
	}

	responseMessage = string(responseBody)
	return responseMessage, nil
}

func (a *ASClient) DeleteResource(repositoryId int, resourceId int) (string, error) {
	responseMessage := ""
	endpoint := fmt.Sprintf("/repositories/%d/resources/%d", repositoryId, resourceId)

	response, err := a.delete(endpoint)
	if err != nil {
		return responseMessage, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return responseMessage, err
	}

	responseMessage = string(responseBody)

	return responseMessage, nil
}

func (a *ASClient) GetResourceTree(repositoryId int, resourceId int) (ResourceTree, error) {
	tree := ResourceTree{}
	endpoint := fmt.Sprintf("/repositories/%d/resources/%d/tree", repositoryId, resourceId)
	response, err := a.get(endpoint, true)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return tree, err
	}

	err = json.Unmarshal(body, &tree)
	if err != nil {
		return tree, err
	}
	return tree, nil

}

func (a *ASClient) GetRandomResourceID() (int, int, error) {
	var repositoryID = 0
	var resourceID = 0
	repositoryIDs, err := a.GetRepositories()
	if err != nil {
		return repositoryID, resourceID, err
	}

	repositoryID = repositoryIDs[rGen.Intn(len(repositoryIDs))]

	resourceIDs, err := a.GetResourceIDs(repositoryID)
	if err != nil {
		return repositoryID, resourceID, err
	}

	resourceID = resourceIDs[rGen.Intn(len(resourceIDs))]

	return repositoryID, resourceID, nil
}

