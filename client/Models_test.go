package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestResourceModel(t *testing.T) {
	var repositoryId, resourceId = 2, 2

	response, err := Client.get(fmt.Sprintf("/repositories/%d/resources/%d", repositoryId, resourceId), true)

	if err != nil {
		t.Error(err)
	}

	body, _ := ioutil.ReadAll(response.Body)

	r := Resource{}
	err = json.Unmarshal(body, &r)
	r.EADID = "XXX"
	if err != nil {
		t.Error(err)
	}

}

func TestResourceModelFail(t *testing.T) {
	var repositoryId, resourceId = 2, 1

	r, err := Client.get(fmt.Sprintf("/repositories/%d/resources/%d", repositoryId, resourceId), true)
	if err == nil {
		t.Error(err)
	}
	r = r

}
