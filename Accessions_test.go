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

}
