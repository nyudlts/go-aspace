package aspace

import (
	"encoding/hex"
	"flag"
	"os"
	"testing"

	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
)

var (
	testClient     *ASClient
	creds          Creds
	testRepoID     = 2
	testResourceID = 1
)

func TestASClient(t *testing.T) {
	flag.Parse()

	t.Run("test get creds", func(t *testing.T) {

		configBytes, err := os.ReadFile(goaspacetest.Config)
		if err != nil {
			t.Error(err)
		}

		creds, err = GetCreds(goaspacetest.Environment, configBytes)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("test ASpace Client Initialization", func(t *testing.T) {
		var err error
		testClient, err = NewClient(goaspacetest.Config, goaspacetest.Environment)
		if err != nil {
			t.Error(err)
		}
		t.Logf("ASpace Client initialized successfully: %s", testClient.RootURL)
	})

	//this is a dumb test
	t.Run("test validate a session key", func(t *testing.T) {
		_, err := hex.DecodeString(testClient.sessionKey)
		if err != nil {
			t.Error(err)
		}

		t.Log("successffully decoded session key")
	})

	t.Run("test get the ASpace server info", func(t *testing.T) {
		var err error
		info, err := testClient.GetAspaceInfo()
		if err != nil {
			t.Error(err)
		}

		t.Logf("ASpace Server Info: %s %s", info.ArchivesSpaceVersion, info.Build)
	})

}
