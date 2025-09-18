package aspace

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/nyudlts/go-aspace/goaspace_testing"
)

func TestRepository(t *testing.T) {
	var (
		repository   *Repository
		repositoryID int
	)
	t.Run("test unmarshal example repository", func(t *testing.T) {
		repoBytes, err := os.ReadFile(filepath.Join(goaspace_testing.TestDataDirJson, "repository.json"))
		if err != nil {
			t.Fatalf("Failed to read repository file: %v", err)
		}

		repository = &Repository{}
		err = json.Unmarshal(repoBytes, repository)
		if err != nil {
			t.Fatalf("Failed to unmarshal repository: %v", err)
		}
		t.Log("Successfully unmarshaled repository:", repository.Name)
	})

	t.Run("test create repository", func(t *testing.T) {
		repository.RepoCode = RandStringRunes(10)
		apiResponse, err := testClient.CreateRepository(repository)
		if err != nil {
			t.Fatal(err)
		}

		if apiResponse.Status != "Created" {
			t.Fatalf("Expected status 'Created', got '%s'", apiResponse.Status)
		}

		repositoryID = apiResponse.ID
		t.Logf("Repository created with ID: %d", apiResponse.ID)
	})

	t.Run("test get repository", func(t *testing.T) {
		repo, err := testClient.GetRepository(repositoryID)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("Retrieved repository: %s", repo.Name)
	})
	t.Run("test update repository", func(t *testing.T) {
		repository.Name = "Updated Repository Name"
		apiResponse, err := testClient.UpdateRepository(repositoryID, repository)
		if err != nil {
			t.Fatal(err)
		}
		if apiResponse.Status != "Updated" {
			t.Fatalf("Expected status 'Updated', got '%s'", apiResponse.Status)
		}

		t.Logf("Repository with ID %d updated successfully", repositoryID)
	})

	t.Run("test get list of repositories", func(t *testing.T) {
		repositories, err := testClient.GetRepositories()
		if err != nil {
			t.Fatal(err)
		}
		if len(repositories) == 0 {
			t.Fatal("Expected at least one repository, got none")
		}

		t.Logf("Retrieved %d repositories", len(repositories))
	})

	t.Run("test get random repository", func(t *testing.T) {
		randomRepoID, err := testClient.GetRandomRepository()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Randomly selected repository ID: %d", randomRepoID)
	})

	t.Run("test delete repository", func(t *testing.T) {
		apiResponse, err := testClient.DeleteRepository(repositoryID)
		if err != nil {
			t.Fatal(err)
		}
		if apiResponse.Status != "Deleted" {
			t.Fatalf("Expected status 'Deleted', got '%s'", apiResponse.Status)
		}
		t.Logf("Repository with ID %d deleted successfully", repositoryID)
	})

}
