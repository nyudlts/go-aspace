package aspace

import (
	"encoding/json"
	"fmt"
	"io"
)

func (a *ASClient) GetRootNode(repositoryID int, resourceID int) (Node, error) {
	Node := Node{}
	endpoint := fmt.Sprintf("/repositories/%d/resources/%d/tree/root", repositoryID, resourceID)
	response, err := a.get(endpoint, true)
	if err != nil {
		return Node, err
	}
	body, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &Node)
	if err != nil {
		return Node, err
	}
	return Node, nil
}

func (a *ASClient) GetNode(repositoryID int, resourceID int, uri string) (Node, error) {
	Node := Node{}
	endpoint := fmt.Sprintf("/repositories/%d/resources/%d/tree/node?node_uri=%s", repositoryID, resourceID, uri)
	response, err := a.get(endpoint, true)
	if err != nil {
		return Node, err
	}
	body, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &Node)
	if err != nil {
		return Node, err
	}
	return Node, nil
}
