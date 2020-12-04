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

	repositoryId, resourceId, err := client.GetRandomResourceID()
	if err != nil {
		t.Error(err)
	}

	resource, err := client.GetResource(repositoryId, resourceId)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("Successfully requested and serialized digital object %s %s\n", resource.URI, resource.Title)
	}

}
