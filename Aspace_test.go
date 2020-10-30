package aspace

import (
	"testing"
)
var RepositoryIDs = []int{2,3,6}
func TestGetAspaceInfo(t *testing.T) {

	info, err := Client.GetAspaceInfo()
	if err != nil {
		t.Error(err)
	}

	t.Log(info)
}

func TestGetResourceIDsByRepository(t *testing.T) {
	repositoryId := 2

	resourceIds, err := Client.GetResourceIDsByRepository(repositoryId)
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
	resources, err := Client.GetResourceIDsByRepository(repositoryId)
	if err != nil {
		t.Error(err)
	}

	resourceId := resources[RandInt(0, len(resources))]

	resource, err := Client.GetResourceByID(repositoryId, resourceId)
	if err != nil {
		t.Error(err)
	}

	title := resource.Title
	t.Log(title)
	if len(title) < 0 {
		t.Errorf("Nil title returned")
	}

}

func TestASClient_GetArchivalObjectById(t *testing.T) {
	ao, err := Client.GetArchivalObjectById(6, 708355)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v\n", ao)
	}

}

func TestASClient_Search(t *testing.T) {
	sr, err := Client.Search(2, "resource", "Donald", 1)
	if err != nil {
		t.Error(err)
	}

	if len(sr.Results) < 1 {
		t.Error(err)
	}
}

func TestASClient_GetDigitalObjectsByRepositoryIdS(t *testing.T) {
	daoIds, err := Client.GetDigitalObjectsByRepositoryId(2)
	if err != nil {
		t.Error(err)
	}
	if len(daoIds) < 1 {
		t.Error("ArchivesSpace returned an empty set")
	}
}

func TestASClient_GetDigitalObject(t *testing.T) {
	_, err := Client.GetDigitalObject(2, 9049)
	if err != nil {
		t.Error(err)
	}
}

