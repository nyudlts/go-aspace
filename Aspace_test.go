package go_aspace

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestGetAspaceInfo(t *testing.T) {

	info, err := GoAspace.GetAspaceInfo()
	if err != nil {
		t.Error(err)
	}

	t.Log(info)
}

func TestGetResourceIDsByRepository(t *testing.T) {
	repositoryId := 2

	resourceIds, err := GoAspace.GetResourceIDsByRepository(repositoryId)
	if err != nil {
		t.Error(err)
	}
	if len(resourceIds) < 0 {
		t.Errorf("There are no resources in repository %d", repositoryId)
	}
}

func TestGetResourceByID(t *testing.T) {

	Seed()
	repositoryId := RepositoryIDs[RandInt(0, 2)]
	resources, err := GoAspace.GetResourceIDsByRepository(repositoryId)
	if err != nil {
		t.Error(err)
	}

	resourceId := resources[RandInt(0, len(resources))]

	resource, err := GoAspace.GetResourceByID(repositoryId, resourceId)
	if err != nil {
		t.Error(err)
	}

	title := resource.Title
	t.Log(title)
	if len(title) < 0 {
		t.Errorf("Nil title returned")
	}

}

func TestASClient_PostResource(t *testing.T) {
	resource, err := GoAspace.GetResourceByID(2, 68)
	if err != nil {
		t.Error(err)
	}
	//t.Logf("%v\n", resource)
	resource.EADID = "zzz"
	j, err := json.MarshalIndent(resource, "", " ")
	p, err := GoAspace.PostResource(2, 68, string(j))
	if err != nil {
		t.Error(err)
	}
	pbody, err := ioutil.ReadAll(p.Body)
	if err != nil {
		t.Error(err)
	}
	r := make(map[string]interface{})
	json.Unmarshal(pbody, &r)
	if r["status"] != "Updated" {
		t.Error(string(pbody))
	}
}
