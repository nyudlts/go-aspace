package aspace

import (
	"encoding/json"
	"fmt"
	"io"
)

func (a *ASClient) GetArchivalObjectIDs(repositoryID int) ([]int, error) {
	aoIds := []int{}
	endpoint := fmt.Sprintf("/repositories/%d/archival_objects?all_ids=true", repositoryID)
	response, err := a.get(endpoint, true)
	if err != nil {
		return aoIds, err
	}
	body, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &aoIds)
	if err != nil {
		return aoIds, err
	}

	return aoIds, nil
}

func (a *ASClient) GetArchivalObject(repositoryId int, aoId int) (ArchivalObject, error) {

	aoURI := fmt.Sprintf("/repositories/%d/archival_objects/%d", repositoryId, aoId)

	return a.GetArchivalObjectFromURI(aoURI)
}

func (a *ASClient) GetArchivalObjectFromURI(aoURI string) (ArchivalObject, error) {

	ao := ArchivalObject{}

	reponse, err := a.get(aoURI, true)
	if err != nil {
		return ao, err
	}

	body, err := io.ReadAll(reponse.Body)
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

	rootNode, err := a.GetRootNode(repositoryId, resourceId)
	if err != nil {
		return aos, err
	}

	aos = []string{}

	for _, node := range rootNode.PrecomputedWaypoints[""].Nodes {
		aos = append(aos, node.URI)
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

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return responseMessage, err
	}

	responseMessage = string(responseBody)
	return responseMessage, nil
}

// func getChildArchivalObjectURIs(children []ResourceTree, aos *[]string) {
// 	for _, child := range children {
// 		*aos = append(*aos, child.RecordURI)
// 		if child.HasChildren {
// 			getChildArchivalObjectURIs(child.Children, aos)
// 		}
// 	}
// }

func (a ASClient) DeleteArchivalObject(repositoryID int, archivalObjectID int) (string, error) {
	responseMessage := ""
	endpoint := fmt.Sprintf("/repositories/%d/archival_objects/%d", repositoryID, archivalObjectID)

	response, err := a.delete(endpoint)
	if err != nil {
		return responseMessage, err
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return responseMessage, err
	}

	responseMessage = string(responseBody)

	return responseMessage, nil
}

// func getChildArchivalObjectURIsFiltered(children []ResourceTree, aos *[]string, filter string) {
// 	matcher := regexp.MustCompile(filter)
// 	for _, child := range children {
// 		if matcher.MatchString(strings.Join(child.InstanceTypes, " ")) == true {
// 			*aos = append(*aos, child.RecordURI)
// 		}

// 		if child.HasChildren {
// 			getChildArchivalObjectURIs(child.Children, aos)
// 		}
// 	}
// }

func (a *ASClient) GetRandomArchivalObject(repositoryID int, resourceID int) (int, int, error) {

	aoURIs, err := a.GetArchivalObjectsForResource(repositoryID, resourceID)
	if err != nil {
		return repositoryID, 0, err
	}

	aoURI := aoURIs[rGen.Intn(len(aoURIs))]
	aspaceURI, err := ParseAspaceURI(aoURI)
	if err != nil {
		return repositoryID, 0, err
	}
	return repositoryID, aspaceURI.ObjectID, nil
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
