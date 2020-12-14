package aspace

import (
	"flag"
	"testing"
)

func TestAccessions(t *testing.T) {
	flag.Parse()
	client, err := NewClient(*envPtr, 10)
	if err != nil {
		t.Error(err)
	}

	t.Run("Test an accession", func(t *testing.T) {
		repositoryID, err := client.GetRandomRepository()
		if err != nil {
			t.Error(err)
		}

		accessionIDs, err := client.GetAccessionIDs(repositoryID)
		if err != nil {
			t.Error(err)
		}

		accessionID := accessionIDs[rGen.Intn(len(accessionIDs))]

		accession, err := client.GetAccession(repositoryID, accessionID)
		if err != nil {
			t.Error(err)
		}

		t.Logf("Successfully serialized %s %s", accession.URI, accession.Title)

	})

}
