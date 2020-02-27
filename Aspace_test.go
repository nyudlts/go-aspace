package go_aspace

import (
	"testing"
)

func TestGetAspaceInfo(t *testing.T) {
	client, err := NewClient(10)
	if err != nil { t.Error(err)}
	GetAspaceInfo(client)

}
