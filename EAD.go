package aspace

import (
	"fmt"
	"github.com/lestrrat-go/libxml2/xsd"
	"github.com/nyudlts/go-aspace/box"
	"io/ioutil"
)

func (a *ASClient) GetEADAsByteArray(repositoryId int, resourceId int) ([]byte, error) {
	eadBytes := []byte{}
	endpoint := fmt.Sprintf("/repositories/%d/resource_descriptions/%d.xml?include_unpublished=%t&include_daos=%t&numbered_cs=%t&ead3=%t&print_pdf=%t", repositoryId, resourceId, false, true, false, false, false)
	response, err := a.get(endpoint, true)
	if err != nil {
		return eadBytes, err
	}

	eadBytes, err = ioutil.ReadAll(response.Body)
	return eadBytes, err
}

func (a *ASClient) SerializeEAD(repositoryId int, resourceId int, daos bool, unpub bool, num_cs bool, ead3 bool, pdf bool) ([]byte, error) {
	var ead []byte = []byte{}

	endpoint := fmt.Sprintf("/repositories/%d/resource_descriptions/%d.xml?include_unpublished=%t&include_daos=%t&numbered_cs=%t&ead3=%t&print_pdf=%t", repositoryId, resourceId, unpub, daos, num_cs, ead3, pdf)
	response, err := a.get(endpoint, true)
	if err != nil {
		return ead, err
	}

	ead, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return ead, err
	}

	return ead, nil

}

func ValidateEAD(fa []byte) error {
	eadxsd, err := xsd.Parse(box.Box.Get("/ead.xsd"))
	if err != nil {
		return err
	}
	doc, err := p.Parse(fa)
	if err != nil {
		return err
	}
	if err := eadxsd.Validate(doc); err != nil {
		return err
	}
	return nil
}
