package aspace

import (
	"flag"
	"testing"
)

var envPtr = flag.String("environment", "dev", "The environment to run tests on")


func TestASClient(t *testing.T) {
	flag.Parse()


	t.Run("Test get session key request", func(t *testing.T) {
		client, err := NewClient(*envPtr, 10)
		if err != nil {
			t.Error(err)
		}
		want := 64
		t.Log(client.sessionKey)
		got := len(client.sessionKey)
		if want != got {
			t.Errorf("wanted key length of %d, got %d", want, got)
		}
	})

}
