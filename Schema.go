package aspace

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func (a *ASClient) GetSchemas() (map[string]interface{}, error) {

	schemas := map[string]interface{}{}
	endpoint := "/schemas"
	response, err := a.get(endpoint, true)
	if err != nil {
		return schemas, err
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return schemas, err
	}

	err = json.Unmarshal(body, &schemas)
	if err != nil {
		return schemas, err
	}

	return schemas, nil
}

func (a *ASClient) GetSchema(s string) ([]byte, error) {

	endpoint := fmt.Sprintf("/schemas/%s", s)
	response, err := a.get(endpoint, true)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	schema, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return schema, nil

}

func (a *ASClient) WriteSchemaFileToDir(schemaName string, path string, perms int) error {
	schemaFile := filepath.Join(path, schemaName+".json")
	schema, err := a.GetSchema(schemaName)
	if err != nil {
		return fmt.Errorf("Error getting schema %s: %v", schemaName, err)
	}
	if err := os.WriteFile(schemaFile, schema, os.FileMode(perms)); err != nil {
		return fmt.Errorf("Error writing schema %s: %v", schemaName, err)
	}

	return nil
}

func (a *ASClient) WriteAllSchemasToDir(path string, perms int) error {
	schemas, err := a.GetSchemas()
	if err != nil {
		return fmt.Errorf("Error getting schemas: %v", err)
	}

	for schemaName := range schemas {
		if err := a.WriteSchemaFileToDir(schemaName, path, perms); err != nil {
			return fmt.Errorf("Error writing schema %s: %v", schemaName, err)
		}
	}

	return nil
}
