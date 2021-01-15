package aspace

import (
	"flag"
	"testing"
)

func TestArchivalObject(t *testing.T) {
	flag.Parse()
	client, err := NewClient(*envPtr, 10)
	if err != nil {
		t.Error(err)
	}

	t.Run("Test serialize and archival object", func(t *testing.T) {
		repositoryID, aoID, err := client.GetRandomArchivalObject()

		ao, err := client.GetArchivalObject(repositoryID, aoID)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Successfully requested and serialized archival %s: %s\n", ao.URI, ao.Title)
		}
	})

	t.Run("Test Basic Search", func(t *testing.T) {

		aos, err := client.SearchArchivalObjects(2, "Broadway")
		if err != nil {
			t.Error(err)
		}

		t.Logf("Search returned %d results", len(aos))
	})

}
