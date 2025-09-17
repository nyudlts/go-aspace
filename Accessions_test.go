package aspace

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/nyudlts/go-aspace/goaspace_testing"
)

var (
	testAccessionIDs = []AccessionEntry{}
	testAccessionID  int
	testAccession    *Accession
)

func TestAccessions(t *testing.T) {
	t.Run("test unmarshal an accession", func(t *testing.T) {

		accessionBytes, err := os.ReadFile(filepath.Join(goaspace_testing.TestDataDir, "accession.json"))
		if err != nil {
			t.Fatal(err)
		}

		testAccession = &Accession{}
		if err := json.Unmarshal(accessionBytes, testAccession); err != nil {
			t.Fatal(err)
		}

		t.Logf("Successfully unmarshalled test accession: %s", testAccession.Title)
	})

	t.Run("Test create an accession", func(t *testing.T) {
		testAccession.ID0 = RandStringRunes(4)
		apiResponse, err := testClient.CreateAccession(testRepoID, *testAccession)
		if err != nil {
			t.Fatal(err)
		}

		testAccessionID = apiResponse.ID

		t.Logf("Successfully created accession: %s", apiResponse.URI)
	})

	/*
		t.Run("Test relate accession to resource", func(t *testing.T) {
			resource, err := testClient.GetResource(testRepoID, testResourceID)
			if err != nil {
				t.Fatalf("Failed to get resource for related accession: %v", err)
			}

			resource.RelatedAccessions = append(resource.RelatedAccessions, RelatedAccession{Ref: testAccession.URI})

			apiResponse, err := testClient.UpdateResource(testRepoID, testResourceID, resource)
			if err != nil {
				t.Fatalf("Failed to update resource with related accession: %v", err)
			}

			t.Logf("Successfully related accession %d to resource %d: %s", testAccessionID, testResourceID, apiResponse.URI)
		})

		t.Run("Test get Accession List for Resource", func(t *testing.T) {
			var err error
			testRepoID, testResourceID, err = testClient.GetRandomResourceID()
			if err != nil {
				t.Fatalf("Failed to get random resource ID: %v", err)
			}

			t.Logf("Testing GetAccessionList for Repository ID %d, ResourceID %d", testRepoID, testResourceID)
			testAccessionIDs, err = testClient.GetAccessionList(testRepoID, testResourceID)
			if err != nil {
				t.Fatalf("Failed to get accession list: %v", err)
			}

			if len(testAccessionIDs) == 0 {
				t.Fatal("Expected at least one accession in the list, but got none")
			}

		})
	*/

	t.Run("Test get an accession", func(t *testing.T) {

		acc, err := testClient.GetAccession(testRepoID, testAccessionID)
		if err != nil {
			t.Fatal(err)
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
			t.Fatal(err)
		}

		if apiResponse.Status != "Updated" {
			t.Fatalf("Expected status 'Updated', got '%s'", apiResponse.Status)
		}

	})

	t.Run("Test delete an accession", func(t *testing.T) {

		apiResponse, err := testClient.DeleteAccession(testRepoID, testAccessionID)
		if err != nil {
			t.Fatal(err)
		}

		if apiResponse.Status != "Deleted" {
			t.Fatalf("Expected status 'Deleted', got '%s'", apiResponse.Status)
		}
	})
}
