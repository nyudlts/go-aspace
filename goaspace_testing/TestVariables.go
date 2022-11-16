package goaspace_testing

import "flag"

var (
	Config      string
	Environment string
)

func init() {
	flag.StringVar(&Config, "config", "", "")
	flag.StringVar(&Environment, "environment", "", "")
}
