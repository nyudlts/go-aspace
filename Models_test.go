package go_aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestResourceModel(t *testing.T) {
	var repositoryId, resourceId = 2, 2312

	a, err := NewClient(10)
	if err != nil {
		t.Error(err)
	}

	resource, err := a.get(fmt.Sprintf("/repositories/%d/resources/%d", repositoryId, resourceId), true)
	if err != nil {
		t.Error(err)
	}

	r := Resource{}
	body, _ := ioutil.ReadAll(resource.Body)
	log.Println(string(body))
	err = json.Unmarshal(body, &r)

	if err != nil {
		t.Error(err)
	}
}

func TestResourceModelFail(t *testing.T) {
	var repositoryId, resourceId = 2, 1

	a, err := NewClient(10)
	if err != nil {
		t.Error(err)
	}

	resource, err := a.get(fmt.Sprintf("/repositories/%d/resources/%d", repositoryId, resourceId), true)
	if err != nil {
		t.Error(err)
	}

	r := Resource{}
	body, _ := ioutil.ReadAll(resource.Body)
	err = json.Unmarshal(body, &r)

	if err != nil {
		t.Error(err)
	}
}
