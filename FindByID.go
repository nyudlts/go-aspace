package aspace

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

type FindArchivalObjectsByIDRefs struct {
	Ref string `json:"ref"`
}
type FindArchivalObjectsByIDResults struct {
	ArchivalObjects []FindArchivalObjectsByIDRefs `json:"archival_objects"`
}

// FindArchivalObjectsByID returns a slice of archival_object_uri strings if
// the id is found, and an empty slice if the id is not found.
//
// Acceptable idTypes are "refID" and "componentID".
//
// The id string will be automatically URL encoded but will not be otherwise transformed
// e.g., ID: 231.0176 -->  ID%3A%20231.0176
//
// This code uses the ASpace "find_by_id" API endpoint described here:
// https://archivesspace.github.io/archivesspace/api/#find-archival-objects-by-ref_id-or-component_id
func (a *ASClient) FindArchivalObjectsByID(repositoryId int, id string, idType string) ([]string, error) {

	aoURIs := []string{}
	// The URL encoding is necessary because the id string will be passed as a query parameter
	id = url.QueryEscape(id)

	idTypeFragment := ""
	switch idType {
	case "refID":
		idTypeFragment = "ref_id"
	case "componentID":
		idTypeFragment = "component_id"
	default:
		return aoURIs, fmt.Errorf("idType must be 'refID' or 'componentID'")
	}

	// Example URL from the documentation:
	// /repositories/:repo_id:/find_by_id/archival_objects?component_id[]=hello_im_a_component_id;resolve[]=archival_objects
	endpoint := fmt.Sprintf(`/repositories/%d/find_by_id/archival_objects?%s[]=%s`, repositoryId, idTypeFragment, id)
	response, err := a.get(endpoint, true)
	if err != nil {
		return aoURIs, err
	}
	body := response.Body
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return aoURIs, err
	}

	response.Body.Close()
	results := FindArchivalObjectsByIDResults{}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return aoURIs, err
	}

	for _, ao := range results.ArchivalObjects {
		aoURIs = append(aoURIs, ao.Ref)
	}
	return aoURIs, nil
}
