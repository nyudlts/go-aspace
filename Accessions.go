package aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (a *ASClient) GetAccessionIDs(repositoryID int) ([]int, error) {
	var accessions = []int{}
	endpoint := fmt.Sprintf("/repositories/%d/accessions?all_ids=true", repositoryID)
	response, err := a.get(endpoint, true)
	if err != nil {
		return accessions, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return accessions, err
	}
	err = json.Unmarshal(body, &accessions)
	if err != nil {
		return accessions, err
	}

	return accessions, nil
}

func (a *ASClient) GetAccession(repositoryID int, accessionID int) (Accession, error) {
	var accession = Accession{}
	endpoint := fmt.Sprintf("/repositories/%d/accessions/%d", repositoryID, accessionID)
	response, err := a.get(endpoint, true)
	if err != nil {
		return accession, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return accession, err
	}
	err = json.Unmarshal(body, &accession)
	if err != nil {
		return accession, err
	}

	return accession, nil
}

func (a *ASClient) UpdateAccession(repositoryID int, accessionID int, accession Accession) (string, error) {
	endpoint := fmt.Sprintf("/repositories/%d/accessions/%d", repositoryID, accessionID)
	body, err := json.Marshal(accession)
	if err != nil {
		return "", err
	}
	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return "", err
	}

	msg, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(msg), nil
}

func (a *ASClient) CreateAccession(repositoryID int, accession Accession) (string, error) {
	endpoint := fmt.Sprintf("/repositories/%d/accessions", repositoryID, accession)
	body, err := json.Marshal(accession)
	if err != nil {
		return "", err
	}
	response, err := a.post(endpoint, true, string(body))
	if err != nil {
		return "", err
	}

	msg, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(msg), nil
}
