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

	r := Resource{}
	err = json.Unmarshal(body, &r)
	r.EADID = "XXX"
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
