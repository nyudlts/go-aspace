package aspace

import (
	"flag"
	"testing"
)


func TestTopContainers(t *testing.T) {
	flag.Parse()
	client, err := NewClient(*envPtr, 10)
	if err != nil {
		t.Error(err)
	}
	t.Run("Test Get TopContainer IDS", func(t *testing.T) {
		topContainers, err := client.GetTopContainerIDs(6)
		if err != nil {
			t.Error(err)
		}

		if len(topContainers) <= 0 {
			t.Error("Array of less than 1 returned")
		} else {
			t.Log("returned", len(topContainers), "Top Containers")
		}

	})

}