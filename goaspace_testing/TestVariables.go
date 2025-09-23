package goaspace_testing

import (
	"flag"
	"path/filepath"
)

var (
	Config            string
	Environment       string
	TestDataDirJson   string = filepath.Join("goaspace_testing", "testdata", "json")
	TestDataDirSchema string = filepath.Join("goaspace_testing", "testdata", "schemas")
	TestDataDirXml    string = filepath.Join("goaspace_testing", "testdata", "xml")
)

func init() {
	flag.StringVar(&Config, "config", "", "")
	flag.StringVar(&Environment, "environment", "", "")
}
