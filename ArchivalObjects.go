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
