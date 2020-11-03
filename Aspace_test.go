package aspace

import (
	"flag"
	"fmt"
	"testing"
)

var RepositoryIDs = []int{2,3,6}

func TestLibrary(t *testing.T) {
	flag.Parse()
	client, err := NewClient(*envPtr, 10)
	if err != nil {
		t.Error(err)
	}

	t.Run("Test Get the ASpace server info", func(t *testing.T) {
		info, err := client.GetAspaceInfo()
		if err != nil {
			t.Error(err)
		}

		t.Log(info)
	})

	t.Run("Test Serialize EAD File", func(t *testing.T) {
		repositoryID := 2
		resourceID := 11
		err := client.SerializeEAD(repositoryID, resourceID, ".", true, false, false, false, false, "test")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Test GET digital object from a repository", func (t *testing.T) {
		daoIds, err := client.GetDigitalObjectsByRepositoryId(2)
		if err != nil {
			t.Error(err)
		}
		if len(daoIds) < 1 {
			t.Error("ArchivesSpace returned an empty set")
		}

		do, err := client.GetDigitalObject(2, daoIds[0])
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("%v\n", do)
	})

}
/*
	t.Run("Test POST digital object", func (t *testing.T){
		do := DigitalObject{
			LockVersion:       0,
			DigitalObjectID:   uuid.Must(uuid.NewRandom()).String(),
			Title:             "Title",
			Publish:           false,
			Restrictions:      false,
			Supressed:         false,
			IsSlugAuto:        false,
			JSONModelType:     "",
			ExternalIds:       nil,
			Subjects:          nil,
			LinkedEvents:      nil,
			Extents:           nil,
			LangMaterials:     nil,
			Dates:             nil,
			ExternalDocuments: nil,
			RightsStatememts:  nil,
			LinkedAgents:      nil,
			FileVersions:      nil,
			Notes:             nil,
			LinkedInstances:   nil,
			URI:               "",
			Repository:        nil,
			Tree:              nil,
		}

		response, err := client.CreateDigitalObject(2, do)
		if err != nil {
			t.Error(err)
		}

		t.Log(response)
	})

	t.Run("Test Deletion of Digital Object", func(t *testing.T) {
		repositoryId := 2
		doId := 41284

		response, err := client.DeleteDigitalObject(repositoryId, doId)
		if err != nil {
			t.Error(err)
		}

		t.Log(response)
	})

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
*/


