package aspace

import (
	"flag"
	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
	"testing"
)

func TestResource(t *testing.T) {
	flag.Parse()
	client, err := NewClient(goaspacetest.ConfigFile, *goaspacetest.EnvPtr, 10)
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
	}

	t.Logf("Successfully requested and serialized %s: %s", resource.URI, resource.Title)

}
