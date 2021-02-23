package aspace

import (
	"flag"
	"fmt"
	"testing"
)

func TestDigitalObject(t *testing.T) {
	flag.Parse()
	client, err := NewClient(*envPtr, 10)
	if err != nil {
		t.Error(err)
	}

	t.Run("Test serialize a digital object", func(t *testing.T) {
		repositoryID, digitalObjectID, err := client.GetRandomDigitalObject()
		if err != nil {
			t.Error(err)
		}

		do, err := client.GetDigitalObject(repositoryID, digitalObjectID)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(fmt.Sprintf("Successfully requested and serialized digital object %s %s\n", do.URI, do.Title))
		}
	})

	t.Run("Test Unamarshal a digital object with notes", func(t *testing.T) {


		do, err := client.GetDigitalObject(2,261)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(do.Notes)
			t.Log(fmt.Sprintf("Successfully requested and serialized digital object %s %s\n", do.URI, do.Title))
		}
	})
}
