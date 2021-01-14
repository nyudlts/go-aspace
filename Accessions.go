package aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Get a list of Accession IDs for a given Repository
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

//Get an Accession object for a given Repository ID and Accession ID
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

//Get a randomly selected Repository and Accession IDs
//This function will return an error if the Repository selected does not contain any Accession records.
func (a *ASClient) GetRandomAccessionID() (int, int, error) {
	repositoryID, err := a.GetRandomRepository()
	if err != nil {
		return 0, 0, err
	}

	log.Println(repositoryID)
	accessionIDs, err := a.GetAccessionIDs(repositoryID)
	if err != nil {
		return 0, 0, err
	}
	log.Println(len(accessionIDs))
	accessionID := accessionIDs[rGen.Intn(len(accessionIDs))]

	return repositoryID, accessionID, nil
}

//Update an Accession record for a given Repository and Accession ID
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

//Create a new Accession with in a given Repository
func (a *ASClient) CreateAccession(repositoryID int, accession Accession) (string, error) {
	endpoint := fmt.Sprintf("/repositories/%d/accessions", repositoryID)
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

//Delete an Accession within a Repository
func (a *ASClient) DeleteAccession(repositoryID int, accessionID int) (string, error) {
	endpoint := fmt.Sprintf("/repositories/%d/accessions/%d", repositoryID, accessionID)
	response, err := a.delete(endpoint)
	if err != nil {
		return fmt.Sprintf("code %d", response.StatusCode), err
	}
	msg, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Sprintf("code %d", response.StatusCode), err
	}

	return string(msg), nil

}
