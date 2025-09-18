package aspace

import (
	"encoding/json"
	"fmt"
	"io"
)

func (a *ASClient) GetSchemas() (map[string]interface{}, error) {
	fmt.Println("test")
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
