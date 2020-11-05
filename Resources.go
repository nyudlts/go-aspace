package aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (a *ASClient) GetResourceIDsByRepository(repositoryId int) ([]int, error) {
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

func (a *ASClient) GetResourceByID(repositoryId int, resourceId int) (Resource, error) {

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

func (a *ASClient) PostResource(repositoryId int, resourceId int, body string) (*http.Response, error) {
	endpoint := fmt.Sprintf("/repositories/%d/resources/%d", repositoryId, resourceId)
	response, err := a.post(endpoint, true, body)
	if err != nil {
		return response, err
	} else {
		return response, nil
	}
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

