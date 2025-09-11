package aspace

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
)

func (a *ASClient) GetResourceIDs(repositoryId int) ([]int, error) {
	var repositoryIds []int
	endpoint := fmt.Sprintf("/repositories/%d/resources?all_ids=true", repositoryId)
	response, err := a.get(endpoint, true)
	if err != nil {
		return repositoryIds, err
	}
	body, _ := io.ReadAll(response.Body)
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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, err
	}
	r.Json = body
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

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return responseMessage, err
	}

	responseMessage = string(responseBody)
	return responseMessage, nil
}

func (a *ASClient) UpdateResourceJson(repositoryID int, resourceID int, resourceJSON []byte) (int, string, error) {

	responseMessage := ""
	endpoint := fmt.Sprintf("/repositories/%d/resources/%d", repositoryID, resourceID)
	response, err := a.post(endpoint, true, string(resourceJSON))
	if err != nil {
		return response.StatusCode, "", err
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return response.StatusCode, responseMessage, err
	}

	responseMessage = string(responseBody)
	return response.StatusCode, responseMessage, nil

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

	responseBody, err := io.ReadAll(response.Body)
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

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return responseMessage, err
	}

	responseMessage = string(responseBody)

	return responseMessage, nil
}

func (a *ASClient) GetResourceTree(repositoryId int, resourceId int) (ResourceTree, error) {
	tree := ResourceTree{}

	endpoint := fmt.Sprintf("/repositories/%d/resources/%d/tree/root", repositoryId, resourceId)

	response, err := a.get(endpoint, true)
	if err != nil {

		return tree, err
	}

	body, err := io.ReadAll(response.Body)
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
	var repositoryID int
	var resourceID int
	repositoryIDs, err := a.GetRepositories()
	if err != nil {
		return 0, 0, err
	}

	repositoryID = repositoryIDs[rGen.Intn(len(repositoryIDs))]

	resourceIDs, err := a.GetResourceIDs(repositoryID)
	if err != nil {
		return repositoryID, resourceID, err
	}

	resourceID = resourceIDs[rGen.Intn(len(resourceIDs))]

	return repositoryID, resourceID, nil
}

func (r Resource) MergeIDs(delimiter string) string {
	ids := r.ID0
	for _, i := range []string{r.ID1, r.ID2, r.ID3} {
		if i != "" {
			ids = ids + delimiter + i
		}
	}
	return ids
}

type ResourceList struct {
	Page     int                 `json:"this_page"`
	LastPage int                 `json:"last_page"`
	Results  []ResourceListEntry `json:"results"`
}

type ResourceListEntry struct {
	ID           string `json:"id"`
	RepositoryID int    `json:"repository_id"`
	ResourceID   int    `json:"resource_id"`
	EADID        string `json:"ead_id"`
	Title        string `json:"title"`
}

func (a *ASClient) GetResourceList(repositoryID int) ([]ResourceListEntry, error) {
	currentPage := 1
	resourceList, err := a.getResourcePage(repositoryID, currentPage)
	if err != nil {
		return nil, err
	}

	entries, err := processEntryList(resourceList)
	if err != nil {
		return nil, err
	}
	lastPage := resourceList.LastPage
	currentPage++
	for i := currentPage; i <= lastPage; i++ {
		r, err := a.getResourcePage(repositoryID, i)
		if err != nil {
			return nil, err
		}
		resources, err := processEntryList(r)
		if err != nil {
			return nil, err
		}
		entries = append(entries, resources...)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Title < entries[j].Title
	})

	return entries, nil
}

func processEntryList(resourceList *ResourceList) ([]ResourceListEntry, error) {
	resources := []ResourceListEntry{}
	for _, resource := range resourceList.Results {
		aspaceURI, err := ParseAspaceURI(resource.ID)
		if err != nil {
			return nil, err
		}
		resource.RepositoryID = aspaceURI.RepositoryID
		resource.ResourceID = aspaceURI.ObjectID
		resources = append(resources, resource)
	}

	return resources, nil
}

func (a *ASClient) getResourcePage(repositoryID int, page int) (*ResourceList, error) {
	endpoint := fmt.Sprintf("/repositories/%d/search?page=%d&type[]=resource&fields[]=id,ead_id,title", repositoryID, page)
	response, err := a.get(endpoint, true)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	resourceList := ResourceList{}
	json.Unmarshal(body, &resourceList)

	return &resourceList, nil
}
