package aspace

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/nyudlts/go-aspace/goaspace_testing"
)

func TestSubjects(t *testing.T) {

	var (
		subject   *Subject
		subjectID int
	)

	t.Run("test unmarshal subject from example JSON", func(t *testing.T) {
		subjectBin, err := os.ReadFile(filepath.Join(goaspace_testing.TestDataDir, "subject.json"))
		if err != nil {
			t.Fatalf("Failed to read subject JSON file: %v", err)
		}

		if err := json.Unmarshal(subjectBin, &subject); err != nil {
			t.Fatalf("Failed to unmarshal subject JSON: %v", err)
		}

		t.Logf("Successfully unmarshaled subject: %s", subject.Title)

	})

	t.Run("test create subject", func(t *testing.T) {
		apiResponse, err := testClient.CreateSubject(subject)
		if err != nil {
			t.Fatalf("Failed to create subject: %v", err)
		}
		if apiResponse.Status != CREATED {
			t.Fatalf("Expected status %s, got %s", CREATED, apiResponse.Status)
		}
		subjectID = apiResponse.ID
		t.Logf("Successfully created subject with id %d", subjectID)

	})

	t.Run("test get subject", func(t *testing.T) {
		var err error
		subject, err = testClient.GetSubject(subjectID)
		if err != nil {
			t.Fatalf("Failed to get subject with id %d: %v", subjectID, err)
		}
		t.Logf("Successfully retrieved subject with uri %s: %s", subject.URI, subject.Title)
	})

	t.Run("test update a subject", func(t *testing.T) {
		updatedTitle := "Updated Subject Title"
		subject.Title = updatedTitle
		apiResponse, err := testClient.UpdateSubject(subjectID, subject)
		if err != nil {
			t.Fatalf("Failed to update subject with id %d: %v", subjectID, err)
		}

		if apiResponse.Status != UPDATED {
			t.Fatalf("Expected status %s, got %s", UPDATED, apiResponse.Status)
		}

		t.Logf("Successfully updated subject with id %d to new title: %s", subjectID, updatedTitle)
	})

	t.Run("test delete a subject", func(t *testing.T) {
		apiResponse, err := testClient.DeleteSubject(subjectID)
		if err != nil {
			t.Fatalf("Failed to delete subject with id %d: %v", subjectID, err)
		}
		if apiResponse.Status != DELETED {
			t.Fatalf("Expected status %s, got %s", DELETED, apiResponse.Status)
		}

		t.Logf("Successfully deleted subject with id %d", subjectID)
	})
}

/*



	t.Run("test get subject", func(t *testing.T) {
		var err error
		subject, err = testClient.GetSubject(subjectID)
		if err != nil {
			t.Fatalf("Failed to get subject with id %d: %v", subjectID, err)
		}
		t.Logf("Successfully retrieved subject with uri %s: %s", subject.URI, subject.Title)
	})

	t.Run("test update a subject", func(t *testing.T) {
		updatedTitle := "Updated Subject Title"
		subject.Title = updatedTitle
		apiResponse, err := testClient.UpdateSubject(subjectID, subject)
		if err != nil {
			t.Fatalf("Failed to update subject with id %d: %v", subjectID, err)
		}

		if apiResponse.Status != UPDATED {
			t.Fatalf("Expected status %s, got %s", UPDATED, apiResponse.Status)
		}

		t.Logf("Successfully updated subject with id %d to new title: %s", subjectID, updatedTitle)
	})


*/
