package aspace

import (
	"flag"
	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
	"testing"
)

func TestCommon(t *testing.T) {
	flag.Parse()
	client, err := NewClient(goaspacetest.ConfigFile, *goaspacetest.EnvPtr, 10)
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
