package go_aspace

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestGetAspaceInfo(t *testing.T) {

	a, err := NewClient(10)
	if err != nil {
		t.Error(err)
	}

	info, err := a.GetAspaceInfo()
	if err != nil {
		t.Error(err)
	}

	t.Log(info)
}

func TestGetResourceIDsByRepository(t *testing.T) {
	repositoryId := 2

	a, err := NewClient(10)
	if err != nil {
		t.Error(err)
	}

	resourceIds, err := a.GetResourceIDsByRepository(repositoryId)
	if err != nil {
		t.Error(err)
	}
	if len(resourceIds) < 0 {
		t.Errorf("There are no resources in repository %d", repositoryId)
	}
}

func TestGetResourceByID(t *testing.T) {

	a, err := NewClient(10)
	if err != nil {
		t.Error(err)
	}

	repositoryId := 2
	resourceId := 2

	resource, err := a.GetResourceByID(repositoryId, resourceId)
	if err != nil {
		t.Error(err)
	}

	title := resource.Title
	if len(title) < 0 {
		t.Errorf("Nil title returned")
	}

	lockVersion := resource.LockVersion

	if lockVersion < 0 {
		t.Errorf("Malformed lock version")
	}

}

func TestASClient_PostResource(t *testing.T) {
	a, err := NewClient(10)
	if err != nil {
		t.Error(err)
	}

	resource, err := a.GetResourceByID(2, 89)
	if err != nil {
		t.Error(err)
	}
	//t.Logf("%v\n", resource)
	resource.EADID = "zzz"
	j, err := json.MarshalIndent(resource, "", " ")
	p, err := a.PostResource(2, 89, string(j))
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
