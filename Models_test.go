package aspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestModels(t *testing.T) {

	client, err := NewClient(*envPtr, 10)
	if err != nil {
		t.Errorf("Not Able to configure client for %s\n", *envPtr)
	}

	t.Run("Test Resource Model", func (t *testing.T) {
		var repositoryId, resourceId = 2, 2

		response, err := client.get(fmt.Sprintf("/repositories/%d/resources/%d", repositoryId, resourceId), true)

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
	})


    t.Run("Test Resource Model Failure", func(t *testing.T) {
		var repositoryId, resourceId = 2, 10000000
		r, err := client.get(fmt.Sprintf("/repositories/%d/resources/%d", repositoryId, resourceId), true)

		if err == nil {
			t.Error(err)
		}

		r = r
	})

}
