package aspace

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
)

var (
	testClient     *ASClient
	creds          Creds
	testRepoID     int
	testResourceID int
)

func TestRemoveTestData(t *testing.T) {
	t.Run("test remove test xml directory", func(t *testing.T) {
		if _, err := os.Stat(goaspacetest.TestDataDirXml); err == nil {
			if err := os.RemoveAll(goaspacetest.TestDataDirXml); err != nil {
				t.Fatalf("Failed to remove test data directory: %v", err)
			}
		}

		if err := os.Mkdir(goaspacetest.TestDataDirXml, 0755); err != nil {
			t.Fatalf("Failed to recreate test xml data directory: %v", err)
		}

		t.Log("Successfully removed and recreated test xml data directory")
	})

	t.Run("test remove test schemas directory", func(t *testing.T) {
		if _, err := os.Stat(goaspacetest.TestDataDirSchema); err == nil {
			if err := os.RemoveAll(goaspacetest.TestDataDirSchema); err != nil {
				t.Fatalf("Failed to remove test schemas directory: %v", err)
			}
		}

		if err := os.Mkdir(goaspacetest.TestDataDirSchema, 0755); err != nil {
			t.Fatalf("Failed to recreate test xml data directory: %v", err)
		}

		t.Log("Successfully removed and recreated test schema directory")
	})
}

func TestASClient(t *testing.T) {
	flag.Parse()

	t.Run("test get creds", func(t *testing.T) {

		configBytes, err := os.ReadFile(goaspacetest.Config)
		if err != nil {
			t.Fatal(err)
		}

		creds, err = GetCreds(goaspacetest.Environment, configBytes)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("test ASpace Client Initialization", func(t *testing.T) {
		var err error
		testClient, err = NewClient(goaspacetest.Config, goaspacetest.Environment)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("ASpace Client initialized successfully: %s", testClient.RootURL)
	})

	//this is a dumb test
	t.Run("test validate a session key", func(t *testing.T) {
		_, err := hex.DecodeString(testClient.sessionKey)
		if err != nil {
			t.Fatal(err)
		}

		t.Log("successffully decoded session key")
	})

	t.Run("test get the ASpace server info", func(t *testing.T) {
		var err error
		info, err := testClient.GetAspaceInfo()
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("ASpace Server Info: %s %s", info.ArchivesSpaceVersion, info.Build)
	})

	//drop existing repository and resources
	t.Run("test delete existing repository and resources", func(t *testing.T) {
		if err := dropAll(); err != nil {
			t.Fatalf("Failed to drop all repositories and resources: %v", err)
		}
		t.Log("Successfully dropped all repositories and resources")
	})

	//create basic object types
	t.Run("test create a repository for testing", func(t *testing.T) {

		repoBin, err := os.ReadFile(filepath.Join(goaspacetest.TestDataDirJson, "repository.json"))
		if err != nil {
			t.Fatal(err)
		}

		repository := &Repository{}
		if err := json.Unmarshal(repoBin, repository); err != nil {
			t.Fatal(err)
		}

		repository.RepoCode = RandStringRunes(20)

		apiResponse, err := testClient.CreateRepository(repository)
		if err != nil {
			t.Fatal(err)
		}
		if apiResponse.Status != "Created" {
			t.Fatalf("Expected status 'Created', got '%s'", apiResponse.Status)
		}
		testRepoID = apiResponse.ID
		t.Logf("Created repository with ID: %d", apiResponse.ID)

	})

	t.Run("test create a resource for testing", func(t *testing.T) {
		// see if resource 1 exists
		_, err := testClient.GetResource(testRepoID, testResourceID)
		if err != nil {
			resourceBin, err := os.ReadFile(filepath.Join(goaspacetest.TestDataDirJson, "resource.json"))
			if err != nil {
				t.Fatal(err)
			}
			resource := &Resource{}
			if err := json.Unmarshal(resourceBin, resource); err != nil {
				t.Fatal(err)
			}

			apiResponse, err := testClient.CreateResource(testRepoID, resource)
			if err != nil {
				t.Fatal(err)
			}
			if apiResponse.Status != "Created" {
				t.Fatalf("Expected status 'Created', got '%s'", apiResponse.Status)
			}

			testResourceID = apiResponse.ID
			t.Logf("Created resource with ID: %d", apiResponse.ID)

		}
	})

}

func dropAll() error {
	repositoryIDs, err := testClient.GetRepositories()
	if err != nil {
		return err
	}

	for _, i := range repositoryIDs {

		fmt.Println("Deleting repository ID:", i)
		resourceIDs, err := testClient.GetResourceIDs(i)
		if err != nil {
			return err
		}

		for _, resourceID := range resourceIDs {
			fmt.Println("Deleting resource:", resourceID)
			if _, err := testClient.DeleteResource(i, resourceID); err != nil {
				return err
			}
		}

		fmt.Println("All resources deleted successfully.")

		_, err = testClient.DeleteRepository(i)
		if err != nil {
			return err
		}

		fmt.Println("Repository deleted successfully.")

	}
	return nil
}
