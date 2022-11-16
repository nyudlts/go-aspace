package aspace

import (
	"encoding/json"
	"fmt"
	"io"
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

func (a *ASClient) GetSchema(s string) (string, error) {
	var schema []byte
	endpoint := fmt.Sprintf("/schemas/%s", s)
	response, err := a.get(endpoint, true)
	if err != nil {
		return "", err
	}
	schema, err = io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(schema), nil
}
