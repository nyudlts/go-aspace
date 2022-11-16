package aspace

import (
	"flag"
	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
	"testing"
)

func TestAccessions(t *testing.T) {
	flag.Parse()
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment, 20)
	if err != nil {
		t.Error(err)
	}

	t.Run("Test an accession", func(t *testing.T) {

		accessionID, RepoID, err := client.GetRandomAccessionID()
		if err != nil {
			t.Error(err)
		}

		accession, err := client.GetAccession(accessionID, RepoID)
		if err != nil {
			t.Error(err)
		}

		t.Logf("Successfully unmarshalled %s %s", accession.URI, accession.Title)

	})

	t.Run("Test get Accession List for Resource", func(t *testing.T) {
		repositoryID := 3
		resourceID := 1823
		accessions, err := client.GetAccessionList(repositoryID, resourceID)
		if err != nil {
			t.Error(err)
		}

		t.Logf("Found %d accessions", len(accessions))
		t.Logf("%v", accessions[0])
	})

}
