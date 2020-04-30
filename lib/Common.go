package lib

import (
	"github.com/lestrrat/go-libxml2/parser"
	"github.com/lestrrat/go-libxml2/xsd"
)

var eadxsd *xsd.Schema
var p = parser.New()

func init() {
	var err error
	eadxsd, err = xsd.ParseFromFile("ead.xsd")
	if err != nil {
		panic(err)
	}
}

func ValidateEAD(fa []byte) error {
	doc, err := p.Parse(fa)
	if err != nil {
		return err
	}

	if err := eadxsd.Validate(doc); err != nil{
		return err
	}
	return nil
}