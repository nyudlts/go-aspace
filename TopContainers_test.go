package aspace

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/nyudlts/go-aspace/goaspace_testing"
)

var topContainerID int
var repositoryID int
var resourceID int

func TestTopContainers(t *testing.T) {
	var (
		topContainer   *TopContainer
		topContainerID int
	)
	t.Run("test unmarshaling top container", func(t *testing.T) {
		tcBin, err := os.ReadFile(filepath.Join(goaspace_testing.TestDataDir, "top_container.json"))
		if err != nil {
			t.Fatalf("could not read top_container.json: %v", err)
		}

		if err = json.Unmarshal(tcBin, topContainer); err != nil {
			t.Fatalf("could not unmarshal top_container.json: %v", err)
		}
	})

	t.Run("test create top container", func(t *testing.T) {
		apiResponse, err := testClient.CreateTopContainer(testRepoID, topContainer)
		if err != nil {
			t.Fatalf("could not create top container: %v", err)
		}

		if apiResponse.Status != "Created" {
			t.Fatalf("expected status 'Created', got '%s'", apiResponse.Status)
		}
		topContainerID = apiResponse.ID
		t.Logf("Created top container with ID %d", topContainerID)
	})

	t.Run("test get top container", func(t *testing.T) {
		var err error
		topContainer, err = testClient.GetTopContainer(testRepoID, topContainerID)
		if err != nil {
			t.Fatalf("could not get top container: %v", err)
		}

		uri := fmt.Sprintf("/repositories/%d/top_containers/%d", testRepoID, topContainerID)
		if topContainer.URI != uri {
			t.Fatalf("expected URI '%s', got '%s'", uri, topContainer.URI)
		}

		t.Logf("Retrieved top container: %s", topContainer.URI)

	})

	t.Run("test update top container", func(t *testing.T) {
		topContainer.CreatedForCollection = "Updated Collection"
		apiResponse, err := testClient.UpdateTopContainer(testRepoID, topContainerID, topContainer)
		if err != nil {
			t.Fatalf("could not update top container: %v", err)
		}

		if apiResponse.Status != "Updated" {
			t.Fatalf("expected status 'Updated', got '%s'", apiResponse.Status)
		}

		t.Logf("Updated top container with ID %d", topContainerID)
	})
	t.Run("test delete top container", func(t *testing.T) {
		apiResponse, err := testClient.DeleteTopContainer(testRepoID, topContainerID)
		if err != nil {
			t.Fatalf("could not delete top container: %v", err)
		}

		if apiResponse.Status != "Deleted" {
			t.Fatalf("expected status 'Deleted', got '%s'", apiResponse.Status)
		}

		t.Logf("Deleted top container with ID %d", topContainerID)
	})
	/*
		flag.Parse()
		client, err := NewClient(goaspacetest.Config, goaspacetest.Environment)
		if err != nil {
			t.Error(err)
		}

		repositoryID, _ = client.GetRandomRepository()
		t.Log("Testing on repository", repositoryID)
		//resourceIDs, _ := client.GetResourceIDs(repositoryID)
		t.Log("Testing on Resource", resourceID)

		t.Run("Test Get TopContainer IDS", func(t *testing.T) {
			topContainers, err := client.GetTopContainerIDs(repositoryID)
			if err != nil {
				t.Error(err)
			}

			if len(topContainers) <= 0 {
				t.Error("Array of less than 1 returned")
			} else {
				t.Log("returned", len(topContainers), "Top Containers")
			}

			topContainerID = topContainers[0]

		})

		t.Run("Test Get A Top Container", func(t *testing.T) {
			topContainer, err := client.GetTopContainer(repositoryID, topContainerID)
			if err != nil {
				t.Error(err)
			}

			t.Log("Top Container", topContainer.URI, "serialized")
		})

		t.Run("Test Get Top Container Ids for Resource", func(t *testing.T) {
			repositoryID, resourceID, err := client.GetRandomResourceID()
			if err != nil {
				t.Error(err)
			}
			t.Log("testing repository:", repositoryID, "resource:", resourceID)
			topContainers, err := client.GetTopContainerIDsForResource(repositoryID, resourceID)
			if err != nil {
				t.Error(err)
			}
			t.Log(topContainers)
		})
	*/
}
