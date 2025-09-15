package aspace

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

// Get a list of Accession IDs for a given Repository ID
func (a *ASClient) GetAccessionIDs(repositoryID int) ([]int, error) {
	var accessions = []int{}
	endpoint := fmt.Sprintf("/repositories/%d/accessions?all_ids=true", repositoryID)
	response, err := a.get(endpoint, true)
	if err != nil {
		return accessions, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return accessions, err
	}
	err = json.Unmarshal(body, &accessions)
	if err != nil {
		return accessions, err
	}

	return accessions, nil
}

// Get an Accession object for a given Repository ID and Accession ID
func (a *ASClient) GetAccession(repositoryID int, accessionID int) (*Accession, error) {
	var accession = Accession{}
	endpoint := fmt.Sprintf("/repositories/%d/accessions/%d", repositoryID, accessionID)
	response, err := a.get(endpoint, true)
	if err != nil {
		return &accession, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return &accession, err
	}
	err = json.Unmarshal(body, &accession)
	if err != nil {
		return &accession, err
	}
	return &accession, nil
}

// Get a randomly selected Repository and Accession IDs
// This function will return an error if the Repository selected does not contain any Accession records.
func (a *ASClient) GetRandomAccessionID() (int, int, error) {
	repositoryID, err := a.GetRandomRepository()
	if err != nil {
		return 0, 0, err
	}

	accessionIDs, err := a.GetAccessionIDs(repositoryID)
	if err != nil {
		return 0, 0, err
	}

	accessionID := accessionIDs[rGen.Intn(len(accessionIDs))]

	return repositoryID, accessionID, nil
}

// Update an Accession record for a given Repository and Accession ID
func (a *ASClient) UpdateAccession(repositoryID int, accessionID int, accession Accession) (*APIResponse, error) {
	endpoint := fmt.Sprintf("/repositories/%d/accessions/%d", repositoryID, accessionID)
	body, err := json.Marshal(accession)
	if err != nil {
		return nil, err
	}

	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return nil, err
	}

	msg, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	apiResponse := &APIResponse{}
	if err = json.Unmarshal(msg, apiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal API response: %v", err)
	}

	return apiResponse, nil
}

// Create a new Accession with in a given Repository
func (a *ASClient) CreateAccession(repositoryID int, accession Accession) (*APIResponse, error) {
	endpoint := fmt.Sprintf("/repositories/%d/accessions", repositoryID)
	body, err := json.Marshal(accession)
	if err != nil {
		return nil, err
	}

	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	apiResponse := &APIResponse{}
	if err = json.Unmarshal(responseBody, apiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal API response: %v", err)
	}

	return apiResponse, nil
}

// Delete an Accession within a Repository
func (a *ASClient) DeleteAccession(repositoryID int, accessionID int) (*APIResponse, error) {
	endpoint := fmt.Sprintf("/repositories/%d/accessions/%d", repositoryID, accessionID)
	response, err := a.delete(endpoint)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	msg, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	apiResponse := &APIResponse{}
	if err = json.Unmarshal(msg, apiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal API response: %v", err)
	}

	return apiResponse, nil

}

func (a Accession) MergeIDs(delimiter string) string {
	ids := a.ID0
	for _, i := range []string{a.ID1, a.ID2, a.ID3} {
		if i != "" {
			ids = ids + delimiter + i
		}
	}
	return ids
}

func (a Accession) GetParentResourceID() int {
	if len(a.RelatedResources) <= 0 {
		return 0
	}
	resId := strings.Split(a.RelatedResources[0]["ref"], "/")[4]
	i, err := strconv.Atoi(resId)
	if err != nil {
		return 0
	}
	return i
}

type AccessionSearchResult struct {
	Page     int              `json:"this_page"`
	LastPage int              `json:"last_page"`
	Results  []AccessionEntry `json:"results"`
}

type AccessionEntry struct {
	ID            string `json:"id"`
	RespositoryID int    `json:"respository_id"`
	ResourceID    int    `json:"resource_id"`
	AccessionID   int    `json:"accession_id"`
	Title         string `json:"title"`
	Identifier    string `json:"identifier"`
}

func (a *ASClient) GetAccessionList(repositoryID int, resourceID int) ([]AccessionEntry, error) {
	currentPage := 1
	accessionList, err := a.getAccessionPage(repositoryID, resourceID, currentPage)
	if err != nil {
		return nil, err
	}

	entries, err := processAccessionList(accessionList, repositoryID, resourceID)
	if err != nil {
		return nil, err
	}

	currentPage++
	for i := currentPage; i <= accessionList.LastPage; i++ {
		al, err := a.getAccessionPage(repositoryID, resourceID, currentPage)
		if err != nil {
			return nil, err
		}
		accessions, err := processAccessionList(al, repositoryID, resourceID)
		if err != nil {
			return nil, err
		}
		entries = append(entries, accessions...)
	}

	sort.Slice(entries, func(i int, j int) bool {
		return entries[i].Title < entries[j].Title
	})

	return entries, nil
}

func processAccessionList(list *AccessionSearchResult, repoID int, resourceID int) ([]AccessionEntry, error) {
	accessions := []AccessionEntry{}
	for _, accession := range list.Results {
		aspaceURI, err := ParseAspaceURI(accession.ID)
		if err != nil {
			return nil, err
		}
		accession.ResourceID = resourceID
		accession.RespositoryID = repoID
		accession.AccessionID = aspaceURI.ObjectID
		accessions = append(accessions, accession)
	}
	return accessions, nil
}

func (a *ASClient) getAccessionPage(repoID int, resID int, currentPage int) (*AccessionSearchResult, error) {
	endpoint := fmt.Sprintf("/repositories/%d/search?page=%d&page_size=50&type[]=accession&fields[]=id,identifier,title&q=repositories/%d/resources/%d", repoID, currentPage, repoID, resID)
	response, err := a.get(endpoint, true)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	accessionSearchResult := AccessionSearchResult{}
	err = json.Unmarshal(body, &accessionSearchResult)
	if err != nil {
		return nil, err
	}
	return &accessionSearchResult, nil
}
