package aspace

import (
	"flag"
	"testing"

	goaspacetest "github.com/nyudlts/go-aspace/goaspace_testing"
)

func TestSchemas(t *testing.T) {
	flag.Parse()
	client, err := NewClient(goaspacetest.Config, goaspacetest.Environment)
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
		_, err := client.GetSchema(randomSchema)
		if err != nil {
			t.Error(err)
		}

		t.Logf("Successfully got schema for %s\n", randomSchema)
	})

	t.Run("Test Writing Schema to File", func(t *testing.T) {
		if err := client.WriteSchemaFileToDir(randomSchema, "schemas", 0644); err != nil {
			t.Fatalf("Error writing schema %s: %v", randomSchema, err)
		}

		t.Logf("Successfully wrote schema %s to file\n", randomSchema)
	})

	t.Run("Test Writing All Schemas to Dir", func(t *testing.T) {
		if err := client.WriteAllSchemasToDir("schemas", 0644); err != nil {
			t.Fatalf("Error writing all schemas: %v", err)
		}

		t.Logf("Successfully wrote all schemas to schemas directory\n")
	})
}
