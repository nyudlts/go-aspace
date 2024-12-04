package aspace

import (
	"flag"
	"testing"

	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
)

func TestSchemas(t *testing.T) {
	flag.Parse()
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment, 20)
	if err != nil {
		t.Error(err)
	}
	var randomSchema string

	t.Run("Test Getting Schema List", func(t *testing.T) {
		schemas, err := client.GetSchemas()
		if err != nil {
			t.Error(err)
		}

		if len(schemas) <= 0 {
			t.Error("Received a zero length list of schemas")
		}

		t.Log("Successfully got a list of schemas")

		keys := make([]string, 0, len(schemas))
		for k := range schemas {
			keys = append(keys, k)
		}

		randomSchema = keys[rGen.Intn(len(keys))]

	})

	t.Run("Test Getting A Schema", func(t *testing.T) {
		schema, err := client.GetSchema(randomSchema)
		if err != nil {
			t.Error(err)
		}

		t.Logf("Successfully got schema for %s: %v\n", randomSchema, schema)
	})
}
