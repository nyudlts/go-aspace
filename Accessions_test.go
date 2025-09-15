package aspace

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/nyudlts/go-aspace/goaspace_testing"
	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
)

var (
	testRepoID       = 2
	testResourceID   = 1
	testAccessionIDs = []AccessionEntry{}
	testAccessionID  int
	testAccession    *Accession
)

func TestAccessions(t *testing.T) {
	//get a client
	flag.Parse()

	var err error
	testClient, err = NewClient(goaspacetest.Config, goaspacetest.Environment)
	if err != nil {
		t.Error(err)
	}

	t.Run("test unmarshal an accession", func(t *testing.T) {

		accessionBytes, err := os.ReadFile(filepath.Join(goaspace_testing.TestDataDir, "accession.json"))
		if err != nil {
			t.Error(err)
		}

		testAccession = &Accession{}
		if err := json.Unmarshal(accessionBytes, testAccession); err != nil {
			t.Error(err)
		}

		t.Logf("Successfully unmarshalled test accession: %s", testAccession.Title)
	})

	t.Run("Test create an accession", func(t *testing.T) {

		apiResponse, err := testClient.CreateAccession(testRepoID, *testAccession)
		if err != nil {
			t.Error(err)
		}

		testAccessionID = apiResponse.ID

		t.Logf("Successfully created accession: %s", apiResponse.URI)
	})

	//get a list of accessions
	t.Run("Test get Accession List for Resource", func(t *testing.T) {
		testRepoID, testResourceID, err = testClient.GetRandomResourceID()
		if err != nil {
			t.Error(err)
		}

		t.Logf("Testing GetAccessionList for Repository ID %d, ResourceID %d", testRepoID, testResourceID)
		testAccessionIDs, err = testClient.GetAccessionList(testRepoID, testResourceID)
		if err != nil {
			t.Error(err)
		}

		t.Logf("Found %d accessions", len(testAccessionIDs))

	})

	t.Run("Test get an accession", func(t *testing.T) {

		acc, err := testClient.GetAccession(testRepoID, testAccessionID)
		if err != nil {
			t.Error(err)
		}

		t.Logf("Successfully requested and serialized accession %s: %s\n", acc.URI, acc.Title)

	})

	t.Run("Test update an accession", func(t *testing.T) {
		testAccession.ID0 = RandStringRunes(4)
		testAccession.ID1 = RandStringRunes(4)
		testAccession.ID2 = RandStringRunes(4)
		testAccession.ID3 = RandStringRunes(4)

		apiResponse, err := testClient.UpdateAccession(testRepoID, testAccessionID, *testAccession)
		if err != nil {
			t.Error(err)
		}

		if apiResponse.Status != "Updated" {
			t.Errorf("Expected status 'Updated', got '%s'", apiResponse.Status)
		}

	})

	t.Run("Test delete an accession", func(t *testing.T) {

		apiResponse, err := testClient.DeleteAccession(testRepoID, testAccessionID)
		if err != nil {
			t.Error(err)
		}

		if apiResponse.Status != "Deleted" {
			t.Errorf("Expected status 'Deleted', got '%s'", apiResponse.Status)
		}
	})
}
