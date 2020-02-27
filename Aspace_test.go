package go_aspace

import (
	"testing"
)

func TestGetAspaceInfo(t *testing.T) {
	client, err := NewClient(10)
	if err != nil {
		t.Error(err)
	}
	info, err := GetAspaceInfo(client)
	if err != nil {
		t.Error(err)
	}
	t.Log(info)
}
