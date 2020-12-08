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
		repositoryId, resourceId, err := client.GetRandomResourceID()
		if err != nil {
			t.Error(err)
		}

		aoURIs, err := client.GetArchivalObjectsForResource(repositoryId, resourceId)
		aoURI := aoURIs[rGen.Intn(len(aoURIs))]
		_, aoID, _ := URISplit(aoURI)
		ao, err := client.GetArchivalObject(repositoryId, aoID)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("Successfully requested and serialized archival %s: %s\n", ao.URI, ao.Title)
		}
	})

}
