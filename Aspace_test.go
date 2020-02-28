package go_aspace

import (
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
	resources, err := a.GetResourceIDsByRepository(repositoryId)
	if err != nil {
		t.Error(err)
	}

	randomNum := randInt(0, len(resources))
	resourceId := resources[randomNum]

	resource, err := a.GetResourceByID(repositoryId, resourceId)
	if err != nil {
		t.Error(err)
	}

	title := resource["title"].(string)
	t.Logf("Retrieved %s", title)

	lockVersion := resource["lock_version"].(float64)

	if lockVersion < 0 {
		t.Errorf("Malformed lock version")
	}

}
