// This file contains tests for the FindByID function in the aspace package.
// All tests are currently commented out because the tests rely on a running ArchivesSpace instance
// with a known set of data.
package aspace

/*
import (
	"flag"
	"testing"

	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
)

func TestFindByIdRefIDSingleResult(t *testing.T) {
	flag.Parse()
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment, 20)
	if err != nil {
		t.Error(t)
	}

	t.Run("Test known archival object by refID", func(t *testing.T) {
		resp, err := client.FindArchivalObjectsByID(3, "ref9000", "refID")
		if err != nil {
			t.Error(t)
		}

		want := "/repositories/3/archival_objects/566184"
		if len(resp) != 1 {
			t.Errorf("got %d, want 1", len(resp))
		} else {
			if resp[0] != want {
				t.Errorf("got %s, want %s", resp[0], want)
			}
		}
	})
}

func TestFindByIdComponentIDSingleResult(t *testing.T) {
	flag.Parse()
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment, 20)
	if err != nil {
		t.Error(t)
	}

	t.Run("Test known archival object by componentID", func(t *testing.T) {
		resp, err := client.FindArchivalObjectsByID(3, "cuid12772", "componentID")
		if err != nil {
			t.Error(t)
		}

		want := "/repositories/3/archival_objects/566184"
		if len(resp) != 1 {
			t.Errorf("got %d, want 1", len(resp))
		} else {
			if resp[0] != want {
				t.Errorf("got %s, want %s", resp[0], want)
			}
		}
	})
}

func TestFindByIdRefIDMultipleResults(t *testing.T) {
	flag.Parse()
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment, 20)
	if err != nil {
		t.Error(t)
	}

	t.Run("Test known multiple archival objects by refID", func(t *testing.T) {
		resp, err := client.FindArchivalObjectsByID(3, "ref5000", "refID")
		if err != nil {
			t.Error(t)
		}

		want := []string{
			"/repositories/3/archival_objects/470159",
			"/repositories/3/archival_objects/476961",
			"/repositories/3/archival_objects/511924",
			"/repositories/3/archival_objects/518940",
			"/repositories/3/archival_objects/546854",
			"/repositories/3/archival_objects/568398",
		}
		if len(resp) != len(want) {
			t.Errorf("got %d, want %d", len(resp), len(want))
		} else {
			for i := range resp {
				if resp[i] != want[i] {
					t.Errorf("got %s, want %s", resp[i], want[i])
				}
			}
		}
	})
}

func TestFindByIdBadIDType(t *testing.T) {
	flag.Parse()
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment, 20)
	if err != nil {
		t.Error(t)
	}

	t.Run("Test bad IDType", func(t *testing.T) {
		_, err := client.FindArchivalObjectsByID(3, "blah", "potato-muffins")
		if err == nil {
			t.Errorf("expected error, got nil")
		}

		want := "idType must be 'refID' or 'componentID'"
		if err.Error() != want {
			t.Errorf("got %s, want %s", err.Error(), want)
		}
	})
}

func TestFindByComponentIDMediaIDSingleResult(t *testing.T) {

	flag.Parse()
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment, 20)
	if err != nil {
		t.Error(t)
	}

	t.Run("Test known archival object by componentID with media ID string", func(t *testing.T) {
		resp, err := client.FindArchivalObjectsByID(3, "ID: 213.0176", "componentID")
		if err != nil {
			t.Error(t)
		}

		want := "/repositories/3/archival_objects/513548"
		if len(resp) != 1 {
			t.Errorf("got %d, want 1", len(resp))
		} else {
			if resp[0] != want {
				t.Errorf("got %s, want %s", resp[0], want)
			}
		}
	})
}

func TestFindByComponentIDNoResults(t *testing.T) {

	flag.Parse()
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment, 20)
	if err != nil {
		t.Error(t)
	}

	t.Run("Test unknown archival object by componentID", func(t *testing.T) {
		resp, err := client.FindArchivalObjectsByID(3, "63a006bc-c059-49fb-a8b0-1cadc4940305", "componentID")
		if err != nil {
			t.Error(t)
		}

		if len(resp) != 0 {
			t.Errorf("got %d, want 0", len(resp))
		}
	})
}
*/
