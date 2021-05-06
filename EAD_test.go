package aspace

import (
	"io/ioutil"
	"testing"
)

func TestEAD(t *testing.T) {

	t.Run("Test EAD 2002 validation", func(t *testing.T) {
		ead, err := ioutil.ReadFile("goaspace_testing/test-ead2002.xml")
		if (err != nil) {
			t.Error(err)
		}

		err = ValidateEAD(ead);
		if err != nil {
			t.Error(err)
		}
	})
}