package go_aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestResourceModel(t *testing.T) {
	var repositoryId, resourceId = 2, 2

	a, err := NewClient(10)

	if err != nil {
		t.Error(err)
	}

	response, err := a.get(fmt.Sprintf("/repositories/%d/resources/%d", repositoryId, resourceId), true)

	if err != nil {
		t.Error(err)
	}

	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))
	r := Resource{}
	err = json.Unmarshal(body, &r)
	r.EAD_ID = "XXX"
	if err != nil {
		t.Error(err)
	}
	j, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(j))
}

func TestResourceModelFail(t *testing.T) {
	var repositoryId, resourceId = 2, 1

	a, err := NewClient(10)
	if err != nil {
		t.Error(err)
	}

	resource, err := a.get(fmt.Sprintf("/repositories/%d/resources/%d", repositoryId, resourceId), true)
	if err == nil {
		t.Error(err)
	}

	r := Resource{}
	body, _ := ioutil.ReadAll(resource.Body)
	err = json.Unmarshal(body, &r)

	if err != nil {
		t.Error(err)
	}
}
