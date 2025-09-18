package aspace

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEAD(t *testing.T) {
	t.Run("test write ead file for resource", func(t *testing.T) {

		resource, err := testClient.GetResource(testRepoID, testResourceID)
		if err != nil {
			t.Fatalf("GetResource returned error: %v", err)
		}

		filename := resource.MergeIDs("_") + "_ead.xml"
		b, err := testClient.GetEADAsByteArray(testRepoID, testResourceID, false)
		if err != nil {
			t.Fatalf("GetEADAsByteArray returned error: %v", err)
		}

		if err := os.WriteFile(filepath.Join("goaspace_testing", "testdata", "xml", filename), b, 0644); err != nil {
			t.Fatalf("Failed to write EAD file: %v", err)
		}

		t.Logf("EAD file written successfully, size: %d bytes", len(b))
	})

}
