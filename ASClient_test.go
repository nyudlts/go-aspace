package aspace

import (
	"encoding/hex"
	"encoding/json"
	"flag"
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

	//create basic object types

	t.Run("test create a repository for testing", func(t *testing.T) {

		repoBin, err := os.ReadFile(filepath.Join(goaspacetest.TestDataDir, "repository.json"))
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
			resourceBin, err := os.ReadFile(filepath.Join(goaspacetest.TestDataDir, "resource.json"))
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
