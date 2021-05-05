package goaspace_testing

import "flag"

var ConfigFile = "/etc/go-aspace.yml"
var EnvPtr = flag.String("environment", "dev", "The environment to run tests on")
