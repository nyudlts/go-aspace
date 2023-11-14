package aspace

import (
	_ "embed"
	"fmt"
	"io"
)

func (a *ASClient) GetMARCAsByteArray(repositoryId int, resourceId int, unpublished bool) ([]byte, error) {
	marcBytes := []byte{}
	endpoint := fmt.Sprintf("/repositories/%d/resources/marc21/%d.xml?include_unpublished_marc=%t", repositoryId, resourceId, unpublished)

	response, err := a.get(endpoint, true)
	if err != nil {
		return marcBytes, err
	}

	marcBytes, err = io.ReadAll(response.Body)
	return marcBytes, err
}
