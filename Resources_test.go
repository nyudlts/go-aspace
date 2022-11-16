package aspace

import (
	"flag"
	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
	"testing"
)

func TestResource(t *testing.T) {
	flag.Parse()
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment, 20)
	if err != nil {
		t.Error(err)
	}

	t.Run("Test serialize Random Resource", func(t *testing.T) {

		repositoryId, resourceId, err := client.GetRandomResourceID()
		if err != nil {
			t.Error(err)
		}

		resource, err := client.GetResource(repositoryId, resourceId)
		if err != nil {
			t.Error(err)
		}

		t.Logf("Successfully requested and serialized %s: %s", resource.URI, resource.Title)
	})

	t.Run("Test request resource list", func(t *testing.T) {
		entries, err := client.GetResourceList(3)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v", entries[0])
		t.Logf("Returned %d entries", len(entries))
	})
}
