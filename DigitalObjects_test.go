package aspace

import (
	"flag"
	"slices"
	"testing"

	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
)

func TestDigitalObject(t *testing.T) {
	flag.Parse()
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment, 20)
	if err != nil {
		t.Error(err)
	}

	t.Run("Test serialize a digital object", func(t *testing.T) {
		repositoryID, digitalObjectID, err := client.GetRandomDigitalObject()
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		do, err := client.GetDigitalObject(repositoryID, digitalObjectID)
		if err != nil {
			t.Error(err)
			t.FailNow()
		} else {
			t.Logf("Successfully requested and serialized digital object %s %s\n", do.URI, do.Title)
		}

	})

	t.Run("Test serialize a digital object using doURI", func(t *testing.T) {
		doURI := "/repositories/3/digital_objects/45726"
		do, err := client.GetDigitalObjectFromURI(doURI)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		want := "MSS_407_cuid29413B"
		if do.DigitalObjectID != want {
			t.Errorf("Expected %s but got %s", want, do.DigitalObjectID)
			t.FailNow()
		}
		t.Logf("Successfully requested and serialized digital object via doURI %s %s\n", do.URI, do.Title)
	})

	t.Run("Test Unmarshal a digital object with notes", func(t *testing.T) {

		do, err := client.GetDigitalObject(2, 261)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(do.Notes)
			t.Logf("Successfully requested and serialized digital object %s %s\n", do.URI, do.Title)
		}
	})

	t.Run("Test DigitalObjectIDs from an ArchivalObject", func(t *testing.T) {

		aoURI := "/repositories/3/archival_objects/912180"

		got, err := client.GetDigitalObjectIDsForArchivalObjectFromURI(aoURI)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		want := []string{"/repositories/3/digital_objects/45716", "/repositories/3/digital_objects/45726", "/repositories/3/digital_objects/45717"}
		if len(want) != len(got) {
			t.Errorf("expected %d digital objects but got %d", len(want), len(got))
			t.FailNow()
		}

		slices.Sort(want)
		slices.Sort(got)

		for i, w := range want {
			if w != got[i] {
				t.Errorf("expected %s but got %s", w, got[i])
				t.FailNow()
			}
		}
		t.Logf("Successfully retrieved DigitalObjectIDs for archival object: %s\n", aoURI)
	})
}
