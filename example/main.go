package main

import (
	"flag"
	"fmt"
	aspace "github.com/nyudlts/go-aspace"
)

var (
	config      string
	environment string
)

func init() {
	flag.StringVar(&config, "config", "", "")
	flag.StringVar(&environment, "environment", "", "")
}

func main() {
	flag.Parse()
	client, err := aspace.NewClient(config, environment, 20)
	if err != nil {
		panic(err)
	}

	ai, err := client.GetAspaceInfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(ai)

}
