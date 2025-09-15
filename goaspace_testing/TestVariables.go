package goaspace_testing

import (
	"flag"
	"path/filepath"
)

var (
	Config      string
	Environment string
	TestDataDir string = filepath.Join("goaspace_testing", "testdata")
)

func init() {
	flag.StringVar(&Config, "config", "", "")
	flag.StringVar(&Environment, "environment", "", "")
}
