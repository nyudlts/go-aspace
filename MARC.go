package aspace

import (
	_ "embed"
	"fmt"
	"github.com/lestrrat-go/libxml2/xsd"
	"io/ioutil"
)

func (a *ASClient) GetMARCAsByteArray(repositoryId int, resourceId int) ([]byte, error) {
	marcBytes := []byte{}
	endpoint := fmt.Sprintf("/repositories/%d/resources/marc21/%d.xml", repositoryId, resourceId)

	response, err := a.get(endpoint, true)
	if err != nil {
		return marcBytes, err
	}

	marcBytes, err = ioutil.ReadAll(response.Body)
	return marcBytes, err
}

func ValidateMARC(fa []byte) error {
	schema, err := schemas.ReadFile("schema/MARC21slim.xsd")
	if err != nil {
		return err
	}
	eadxsd, err := xsd.Parse(schema)
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
