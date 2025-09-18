package aspace

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMARC(t *testing.T) {
	t.Run("test write marc21 file for resource", func(t *testing.T) {

		resource, err := testClient.GetResource(testRepoID, testResourceID)
		if err != nil {
			t.Fatalf("GetResource returned error: %v", err)
		}

		filename := resource.MergeIDs("_") + "_marc.xml"
		b, err := testClient.GetMARCAsByteArray(testRepoID, testResourceID, false)
		if err != nil {
			t.Fatalf("GetMARCAsByteArray returned error: %v", err)
		}

		if err := os.WriteFile(filepath.Join("goaspace_testing", "testdata", "xml", filename), b, 0644); err != nil {
			t.Fatalf("Failed to write MARCXML file: %v", err)
		}

		t.Logf("MARCXML file written successfully, size: %d bytes", len(b))
	})

}
