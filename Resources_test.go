package aspace

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/nyudlts/go-aspace/goaspace_testing"
)

func TestResource(t *testing.T) {

	var (
		resource   *Resource
		resourceID int
	)

	t.Run("test unmarshal resource json to struct", func(t *testing.T) {
		resourceBin, err := os.ReadFile(filepath.Join(goaspace_testing.TestDataDir, "resource.json"))
		if err != nil {
			t.Fatal(err)
		}

		resource = &Resource{}
		err = json.Unmarshal(resourceBin, resource)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("Successfully unmarshaled resource: %s", resource.Title)
	})

	t.Run("test create resource", func(t *testing.T) {
		resource.ID0 = RandStringRunes(6)
		apiResponse, err := testClient.CreateResource(testRepoID, resource)
		if err != nil {
			t.Fatal(err)
		}

		if apiResponse.Status != CREATED {
			t.Fatalf("Expected status %s, got %s", CREATED, apiResponse.Status)
		}

		resourceID = apiResponse.ID

		t.Logf("Successfully created resource: %d", resourceID)
	})

	t.Run("test get resource by ID", func(t *testing.T) {
		var err error
		resource, err = testClient.GetResource(testRepoID, resourceID)
		if err != nil {
			t.Fatal(err)
		}

		uri := fmt.Sprintf("/repositories/%d/resources/%d", testRepoID, resourceID)
		if resource.URI != uri {
			t.Fatalf("Expected URI %s, got %s", uri, resource.URI)
		}
	})

	/*
		t.Run("test get a random resource", func(t *testing.T) {

			repositoryId, resourceId, err := testClient.GetRandomResourceID()
			if err != nil {
				t.Error(err)
			}

			resource, err := testClient.GetResource(repositoryId, resourceId)
			if err != nil {
				t.Error(err)
			}

			t.Logf("Successfully requested and serialized %s: %s", resource.URI, resource.Title)
		})

		t.Run("get resource ids for repository", func(t *testing.T) {
			entries, err := testClient.GetResourceIDsForRepository(testRepoID)
			if err != nil {
				t.Fatal(err)
			}

			if len(entries) == 0 {
				t.Fatal("Expected to get at least one resource ID, got none")
			}

			t.Logf("Successfully retrieved %d resource IDs for repository %d", len(entries), testRepoID)
		})
	*/

	t.Run("test update resource", func(t *testing.T) {

		updatedTitle := "Updated Resource Title"
		resource.Title = updatedTitle
		apiResponse, err := testClient.UpdateResource(testRepoID, resourceID, resource)
		if err != nil {
			t.Fatal(err)
		}
		if apiResponse.Status != UPDATED {
			t.Fatalf("Expected status %s, got %s", UPDATED, apiResponse.Status)
		}

		t.Logf("Successfully updated resource: %s", resource.URI)
	})

	t.Run("test delete resource", func(t *testing.T) {
		apiResponse, err := testClient.DeleteResource(testRepoID, resourceID)
		if err != nil {
			t.Fatal(err)
		}
		if apiResponse.Status != DELETED {
			t.Fatalf("Expected status %s, got %s", DELETED, apiResponse.Status)
		}
		t.Logf("Successfully deleted resource with ID: %d", resourceID)
	})

}
