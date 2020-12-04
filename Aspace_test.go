package aspace

import (
	"flag"
	"testing"
)

var client *ASClient

func TestLibrary(t *testing.T) {

	flag.Parse()
	client, err := NewClient(*envPtr, 10)
	if err != nil {
		t.Error(err)
	}

	t.Run("Test Get the ASpace server info", func(t *testing.T) {
		info, err := client.GetAspaceInfo()
		if err != nil {
			t.Error(err)
		}

		t.Log(info)
	})

}
