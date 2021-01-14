package aspace

import (
	"flag"
	"testing"
)

func TestSearch(t *testing.T) {
	flag.Parse()
	client, err := NewClient(*envPtr, 10)
	if err != nil {
		t.Error(t)
	}

	t.Run("Test A Basic Search", func(t *testing.T) {
		resp, err := client.Search(2, "archival_object", "records", 1)
		if err != nil {
			t.Error(t)
		}
		t.Log(resp.Results[0])
	})

}
