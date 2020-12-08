package aspace

import (
	"flag"
	"testing"
)

var envPtr = flag.String("environment", "dev", "The environment to run tests on")

func TestCommon(t *testing.T) {
	flag.Parse()
	client, err := NewClient(*envPtr, 10)
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
}
