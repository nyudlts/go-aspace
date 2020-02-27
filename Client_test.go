package go_aspace

import (
	"testing"
)

func TestNewClientHasSession(t *testing.T) {
	client, err := NewClient(10)
	if err != nil {
		t.Error(err)
	}

	want := 64
	got := len(client.sessionKey)
	if want != got {
		t.Errorf("wanted key length of %d, got %d", want, got)
	}

}
