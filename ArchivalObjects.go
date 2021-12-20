package aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
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

func (a *ASClient) GetArchivalObjectsForResource(repositoryId int, resourceId int, filter string) ([]string, error) {

	aos := []string{}
	tree, err := a.GetResourceTree(repositoryId, resourceId)
	if err != nil {
		return aos, err
	}

	if filter == "" {
		getChildArchivalObjectURIsFiltered(tree.Children, &aos, filter)
	} else {
		getChildArchivalObjectURIs(tree.Children, &aos)
	}
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

func (a ASClient) DeleteArchivalObject(repositoryID int, archivalObjectID int) (string, error) {
	responseMessage := ""
	endpoint := fmt.Sprintf("/repositories/%d/archival_objects/%d", repositoryID, archivalObjectID)

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

func getChildArchivalObjectURIsFiltered(children []ResourceTree, aos *[]string, filter string) {
	matcher := regexp.MustCompile(filter)
	for _, child := range children {
		if matcher.MatchString(strings.Join(child.InstanceTypes, " ")) == true {
			*aos = append(*aos, child.RecordURI)
		}

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
	aoURIs, err := a.GetArchivalObjectsForResource(repositoryID, resourceID, "")
	aoURI := aoURIs[rGen.Intn(len(aoURIs))]
	_, aoID, _ := URISplit(aoURI)
	return repositoryID, aoID, nil
}

func (a *ASClient) SearchArchivalObjects(repoID int, s string) ([]ArchivalObject, error) {
	aos := []ArchivalObject{}
	iresults, err := a.Search(repoID, "archival_object", s, 1)

	if err != nil {
		return aos, err
	}

	lastPage := iresults.LastPage

	//iterate here through pages
	for i := 1; i < lastPage; i++ {
		results, err := a.Search(repoID, "archival_object", s, i)
		for _, r := range results.Results {
			ao_json := []byte(fmt.Sprintf("%v", r["json"]))
			if err != nil {
				return aos, err
			}
			ao := ArchivalObject{}
			err := json.Unmarshal(ao_json, &ao)
			if err != nil {
				return aos, err
			}
			aos = append(aos, ao)
		}
	}

	return aos, nil
}
