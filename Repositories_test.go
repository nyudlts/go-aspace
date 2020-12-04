package aspace

import (
	"flag"
	"testing"
)

func TestRepository(t *testing.T) {
	flag.Parse()
	client, err := NewClient(*envPtr, 10)
	if err != nil {
		t.Error(err)
	}

	repositoryID, err := client.GetRandomRepository()
	if err != nil {
		t.Error(err)
	}

	repository, err := client.GetRepository(repositoryID)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Successfully requested and serialized repository %s: %s\n", repository.URI, repository.Name)
}
