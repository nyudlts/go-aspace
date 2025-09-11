package aspace

import (
	"flag"
	"strconv"
	"testing"

	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
)

func TestCommon(t *testing.T) {
	flag.Parse()
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment)
	if err != nil {
		t.Error(err)
	}

	t.Run("Test get the ASpace server info", func(t *testing.T) {
		info, err := client.GetAspaceInfo()
		if err != nil {
			t.Error(err)
		}

		t.Log(info)
	})

	t.Run("Test get session key", func(t *testing.T) {

		want := 64
		t.Log(client.sessionKey)
		got := len(client.sessionKey)
		if want != got {
			t.Errorf("wanted key length of %d, got %d", want, got)
		} else {
			t.Log("Succesfully requested valid session key.")
		}
	})

	t.Run("Test ParseCreateOrUpdateResponse() Created", func(t *testing.T) {

		responseBody := `{"status":"Created","id":61344,"lock_version":0,"stale":true,"uri":"/repositories/6/digital_objects/61344","warnings":[]}`
		got := ParseCreateOrUpdateResponse(responseBody)

		scenarios := [][]string{
			{"Created", got.Status, "Incorrect Status"},
			{"", got.Error, "Incorrect Error"},
			{"61344", strconv.FormatInt(int64(got.ID), 10), "Incorrect ID"},
			{"0", strconv.FormatInt(int64(got.LockVersion), 10), "Incorrect LockVersion"},
			{"true", strconv.FormatBool(got.Stale), "Incorrect Stale"},
			{"/repositories/6/digital_objects/61344", got.URI, "Incorrect URI"},
		}

		for _, scenario := range scenarios {
			if scenario[1] != scenario[0] {
				t.Errorf("unexpected result: %s: want: '%s', got: '%s'", scenario[2], scenario[0], scenario[1])
			}
		}
	})

	t.Run("Test ParseCreateOrUpdateResponse() Error", func(t *testing.T) {

		responseBody := `{"error":"I need more coffee!"}`
		got := ParseCreateOrUpdateResponse(responseBody)

		scenarios := [][]string{
			{"", got.Status, "Incorrect Status"},
			{"I need more coffee!", got.Error, "Incorrect Error"},
			{"0", strconv.FormatInt(int64(got.ID), 10), "Incorrect ID"},
			{"0", strconv.FormatInt(int64(got.LockVersion), 10), "Incorrect LockVersion"},
			{"false", strconv.FormatBool(got.Stale), "Incorrect Stale"},
			{"", got.URI, "Incorrect URI"},
		}

		for _, scenario := range scenarios {
			if scenario[1] != scenario[0] {
				t.Errorf("unexpected result: %s: want: '%s', got: '%s'", scenario[2], scenario[0], scenario[1])
			}
		}
	})

	t.Run("Test AspaceURI type", func(t *testing.T) {
		uri := "/repositories/3/digital_objects/2167"
		got, err := ParseAspaceURI(uri)
		if err != nil {
			t.Error(err)
		}

		want := AspaceURI{RepositoryID: 3, ObjectType: "digital_objects", ObjectID: 2167}

		if got != want {
			t.Errorf("Got %s, wanted: %s", got, want)
		}

	})

}
