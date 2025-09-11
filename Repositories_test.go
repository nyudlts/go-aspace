package aspace

import (
	"flag"
	"testing"

	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
)

func TestRepository(t *testing.T) {
	flag.Parse()
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment)
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
