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
