package aspace

import (
	"flag"
	"testing"

	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
)

func TestArchivalObject(t *testing.T) {
	flag.Parse()
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment)
	if err != nil {
		t.Error(err)
	}

	t.Run("Test serialize an archival object", func(t *testing.T) {
		repositoryID, resourceID, err := client.GetRandomResourceID()
		if err != nil {
			t.Error(err)
		}

		t.Log("Testing repoID: ", repositoryID, " resourceID: ", resourceID)

		repositoryID, aoID, err := client.GetRandomArchivalObject(repositoryID, resourceID)
		if err != nil {
			t.Error(err)
		}

		ao, err := client.GetArchivalObject(repositoryID, aoID)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Successfully requested and serialized archival object %s: %s\n", ao.URI, ao.Title)
		}
	})

	/*
		t.Run("Test Basic Search", func(t *testing.T) {

			aos, err := client.SearchArchivalObjects(2, "Broadway")
			if err != nil {
				t.Error(err)
			}

			t.Logf("Search returned %d results", len(aos))
		})
	*/
}
