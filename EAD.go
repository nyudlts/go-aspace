package aspace

import (
	"fmt"
	"io"
)

func (a *ASClient) GetEADAsByteArray(repositoryId int, resourceId int, unpublished bool) ([]byte, error) {
	eadBytes, err := a.SerializeEAD(repositoryId, resourceId, true, unpublished, false, false, false)
	if err != nil {
		return nil, err
	}
	return eadBytes, err
}

func (a *ASClient) SerializeEAD(repositoryId int, resourceId int, daos bool, unpub bool, num_cs bool, ead3 bool, pdf bool) ([]byte, error) {
	var ead []byte = []byte{}

	endpoint := fmt.Sprintf("/repositories/%d/resource_descriptions/%d.xml?include_unpublished=%t&include_daos=%t&numbered_cs=%t&ead3=%t&print_pdf=%t", repositoryId, resourceId, unpub, daos, num_cs, ead3, pdf)
	response, err := a.get(endpoint, true)
	if err != nil {
		return ead, err
	}

	ead, err = io.ReadAll(response.Body)
	if err != nil {
		return ead, err
	}

	return ead, nil

}
