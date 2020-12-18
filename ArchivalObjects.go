package aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (a *ASClient) GetArchivalObject(repositoryId int, aoId int) (ArchivalObject, error) {

	ao := ArchivalObject{}
	endpoint := fmt.Sprintf("/repositories/%d/archival_objects/%d", repositoryId, aoId)

	reponse, err := a.get(endpoint, true)
	if err != nil {
		return ao, err
	}

	body, err := ioutil.ReadAll(reponse.Body)
	if err != nil {
		return ao, err
	}

	err = json.Unmarshal(body, &ao)
	if err != nil {
		return ao, err
	}

	return ao, nil
}

func (a *ASClient) GetArchivalObjectsForResource(repositoryId int, resourceId int) ([]string, error) {

	aos := []string{}
	tree, err := a.GetResourceTree(repositoryId, resourceId)
	if err != nil {
		return aos, err
	}

	getChildArchivalObjectURIs(tree.Children, &aos)

	return aos, nil
}

func (a *ASClient) UpdateArchivalObject(repositoryId int, archivalObjectId int, archivalObject ArchivalObject) (string, error) {
	responseMessage := ""
	endpoint := fmt.Sprintf("/repositories/%d/archival_objects/%d", repositoryId, archivalObjectId)
	body, err := json.Marshal(archivalObject)
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

func getChildArchivalObjectURIs(children []ResourceTree, aos *[]string) {
	for _, child := range children {
		*aos = append(*aos, child.RecordURI)
		if child.HasChildren {
			getChildArchivalObjectURIs(child.Children, aos)
		}
	}
}

func (a *ASClient) GetRandomArchivalObject() (int, int, error) {
	repositoryID, resourceID, err := a.GetRandomResourceID()
	if err != nil {
		return 0, 0, err
	}
	aoURIs, err := a.GetArchivalObjectsForResource(repositoryID, resourceID)
	aoURI := aoURIs[rGen.Intn(len(aoURIs))]
	_, aoID, _ := URISplit(aoURI)
	return repositoryID, aoID, nil
}
