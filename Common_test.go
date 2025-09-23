package aspace

import (
	"flag"
	"strconv"
	"testing"
)

func TestCommon(t *testing.T) {
	flag.Parse()

	t.Run("Test ParseAPIResponse() Created", func(t *testing.T) {

		responseBody := `{"status":"Created","id":61344,"lock_version":0,"stale":true,"uri":"/repositories/6/digital_objects/61344","warnings":[]}`
		got := ParseAPIResponse(responseBody)

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

	t.Run("Test ParseAPIResponse() Error", func(t *testing.T) {

		responseBody := `{"error":"I need more coffee!"}`
		got := ParseAPIResponse(responseBody)

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
