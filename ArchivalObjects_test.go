package aspace

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/nyudlts/go-aspace/goaspace_testing"
)

func TestArchivalObject(t *testing.T) {
	var (
		ao   *ArchivalObject
		aoID int
	)

	t.Run("test serialize example archival object", func(t *testing.T) {
		aoBytes, err := os.ReadFile(filepath.Join(goaspace_testing.TestDataDir, "archival_object.json"))
		if err != nil {
			t.Error(err)
		}

		ao = &ArchivalObject{}
		err = json.Unmarshal(aoBytes, ao)
		if err != nil {
			t.Errorf("Error unmarshalling archival object: %v", err)
		}

		t.Logf("Successfully unmarshalled archival object: %s", ao.Title)
	})

	t.Run("test create an archival object", func(t *testing.T) {
		ao.Resource.Ref = fmt.Sprintf("/repositories/%d/resources/%d", testRepoID, testResourceID) // Set the resource reference
		apiResponse, err := testClient.CreateArchivalObject(testRepoID, ao)
		if err != nil {
			t.Error(err)
		}

		if apiResponse.Status != "Created" {
			t.Errorf("Expected status 'Created', got '%s'", apiResponse.Status)
		}

		t.Logf("Successfully created archival object with ID %d: %s", apiResponse.ID, ao.Title)
		aoID = apiResponse.ID
	})

	t.Run("test get an archival object", func(t *testing.T) {
		ao, err := testClient.GetArchivalObject(testRepoID, aoID)
		if err != nil {
			t.Error(err)
		}

		t.Logf("Successfully retrieved archival object %d: %s", aoID, ao.Title)
	})

	t.Run("test update an archival object", func(t *testing.T) {
		ao.Title = "Updated Archival Object Title"
		apiResponse, err := testClient.UpdateArchivalObject(testRepoID, aoID, ao)
		if err != nil {
			t.Error(err)
		}
		if apiResponse.Status != "Updated" {
			t.Errorf("Expected status 'Updated', got '%s'", apiResponse.Status)
		}

		t.Logf("Successfully updated archival object %d: %s", aoID, ao.Title)
	})

	t.Run("test delete an archival object", func(t *testing.T) {
		apiResponse, err := testClient.DeleteArchivalObject(testRepoID, aoID)
		if err != nil {
			t.Error(err)
		}

		if apiResponse.Status != "Deleted" {
			t.Errorf("Expected status 'Deleted', got '%s'", apiResponse.Status)
		}
	})
}

/*

	//this is broken
	t.Run("test basic search", func(t *testing.T) {

		aos, err := testClient.SearchArchivalObjects(2, "Archival Object title")
		if err != nil {
			t.Error(err)
		}

		if len(aos) == 0 {
			t.Error("Expected to find at least one archival object, but found none")
		}
	})

}
*/
