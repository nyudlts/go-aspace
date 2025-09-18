package aspace

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/nyudlts/go-aspace/goaspace_testing"
)

func TestDigitalObject(t *testing.T) {
	var (
		do   *DigitalObject
		doID int
	)

	t.Run("test unmarshal a digital object", func(t *testing.T) {
		doBytes, err := os.ReadFile(goaspace_testing.TestDataDirJson + "/digital_object.json")
		if err != nil {
			t.Error(err)
		}

		do = &DigitalObject{}
		if err := json.Unmarshal(doBytes, do); err != nil {
			t.Error(err)
		} else {
			t.Logf("Successfully unmarshalled digital object: %s", do.Title)
		}
	})

	t.Run("test create a digital object", func(t *testing.T) {
		apiResponse, err := testClient.CreateDigitalObject(testRepoID, do)
		if err != nil {
			t.Error(err)
		} else {

			if apiResponse.Status != "Created" {
				t.Errorf("Expected status 'Created' but got '%s'", apiResponse.Status)
			} else {
				doID = apiResponse.ID
				t.Logf("Successfully created digital object: %s", apiResponse.URI)
			}
		}
	})

	t.Run("test update a digital object", func(t *testing.T) {
		do.Title = "Updated Digital Object Title"
		apiResponse, err := testClient.UpdateDigitalObject(testRepoID, doID, do)
		if err != nil {
			t.Error(err)
		} else {
			if apiResponse.Status != "Updated" {
				t.Errorf("Expected status 'Updated' but got '%s'", apiResponse.Status)
			} else {
				t.Logf("Successfully updated digital object: %s", apiResponse.URI)
			}
		}
	})

	t.Run("test get digital object uris for resource", func(t *testing.T) {
		doURIs, err := testClient.GetDigitalObjectIDs(testRepoID)
		if err != nil {
			t.Error(err)
		} else {
			if len(doURIs) == 0 {
				t.Error("Expected to find digital object URIs but got none")
			} else {
				t.Logf("Found %d digital object URIs for repository %d", len(doURIs), testRepoID)
			}
		}
	})

	t.Run("test delete a digital object", func(t *testing.T) {
		apiResponse, err := testClient.DeleteDigitalObject(testRepoID, doID)
		if err != nil {
			t.Error(err)
		} else {

			if apiResponse.Status != "Deleted" {
				t.Errorf("Expected status 'Deleted' but got '%s'", apiResponse.Status)
			} else {
				t.Logf("Successfully deleted digital object %s", do.URI)
			}
		}
	})

}

/*
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment)
	if err != nil {
		t.Error(err)
	}

	t.Run("Test serialize a digital object", func(t *testing.T) {
		repositoryID, digitalObjectID, err := client.GetRandomDigitalObject()
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		do, err := client.GetDigitalObject(repositoryID, digitalObjectID)
		if err != nil {
			t.Error(err)
			t.FailNow()
		} else {
			t.Logf("Successfully requested and serialized digital object %s %s\n", do.URI, do.Title)
		}

	})

	t.Run("Test serialize a digital object using doURI", func(t *testing.T) {
		doURI := "/repositories/3/digital_objects/45726"
		do, err := client.GetDigitalObjectFromURI(doURI)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		want := "MSS_407_cuid29413B"
		if do.DigitalObjectID != want {
			t.Errorf("Expected %s but got %s", want, do.DigitalObjectID)
			t.FailNow()
		}
		t.Logf("Successfully requested and serialized digital object via doURI %s %s\n", do.URI, do.Title)
	})

	t.Run("Test Unmarshal a digital object with notes", func(t *testing.T) {

		do, err := client.GetDigitalObject(2, 261)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(do.Notes)
			t.Logf("Successfully requested and serialized digital object %s %s\n", do.URI, do.Title)
		}
	})

	t.Run("Test DigitalObjectIDs from an ArchivalObject", func(t *testing.T) {

		aoURI := "/repositories/3/archival_objects/912180"

		got, err := client.GetDigitalObjectIDsForArchivalObjectFromURI(aoURI)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		want := []string{"/repositories/3/digital_objects/45716", "/repositories/3/digital_objects/45726", "/repositories/3/digital_objects/45717"}
		if len(want) != len(got) {
			t.Errorf("expected %d digital objects but got %d", len(want), len(got))
			t.FailNow()
		}

		slices.Sort(want)
		slices.Sort(got)

		for i, w := range want {
			if w != got[i] {
				t.Errorf("expected %s but got %s", w, got[i])
				t.FailNow()
			}
		}
		t.Logf("Successfully retrieved DigitalObjectIDs for archival object: %s\n", aoURI)
	})
}
*/
