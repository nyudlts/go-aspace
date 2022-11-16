package aspace

import (
	"os"
	"testing"
)

func TestEAD(t *testing.T) {

	t.Run("Test EAD 2002 validation", func(t *testing.T) {
		ead, err := os.ReadFile("goaspace_testing/test-ead2002.xml")
		if err != nil {
			t.Error(err)
		}

		err = ValidateEAD(ead)
		if err != nil {
			t.Error(err)
		}
	})
}
