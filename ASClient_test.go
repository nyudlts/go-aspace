package aspace

import (
	"flag"
	"net/http"
	"os"
	"testing"

	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
)

var creds Creds

func TestASClient(t *testing.T) {
	flag.Parse()

	t.Run("Test get creds", func(t *testing.T) {

		configBytes, err := os.ReadFile(goaspacetest.Config)
		if err != nil {
			t.Error(err)
		}

		creds, err = GetCreds(goaspacetest.Environment, configBytes)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Test Endpoint", func(t *testing.T) {
		resp, err := http.Get(creds.URL)
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()
	})

	t.Run("Test ASpace Client Initialization", func(t *testing.T) {
		client, err := NewClient(goaspacetest.Config, goaspacetest.Environment)
		if err != nil {
			t.Error(err)
		}
		t.Logf("ASpace Client initialized successfully: %s", client.RootURL)
	})

}
