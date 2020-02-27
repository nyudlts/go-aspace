package go_aspace

import (
	"testing"
)

func TestGetAspaceInfo(t *testing.T) {

	info, err := GetAspaceInfo()
	if err != nil {
		t.Error(err)
	}
	t.Log(info)
}

func TestGetResourceIDsByRepository(t *testing.T) {
	repositoryId := 2
	resourceIds, err := GetResourceIDsByRepository(repositoryId)
	if err != nil {
		t.Error(err)
	}
	if len(resourceIds) < 0 {
		t.Errorf("There are no resources in repository %d", repositoryId)
	}
}

func TestGetResourceByID(t *testing.T) {
	repositoryId := 2
	resources, err := GetResourceIDsByRepository(repositoryId)
	if err != nil {
		t.Error(err)
	}

	randomNum := randInt(0, len(resources))
	resourceId := resources[randomNum]

	resource, err := GetResourceByID(repositoryId, resourceId)
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
