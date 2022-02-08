# go-aspace v0.3.12b
a Go library for ArchivesSpace integrations

## Use
1. Get the go-aspace library: $go get github.com/nyudlts/go-aspace.
2. Edit the config file `$GOROOT/github.com/nyudlts/go-aspace/go-aspace.yml_template`, enter your aspace credentials
4. Import "github.com/nyudlts/go-aspace" into your project
5. Create an instance of the aspace client:

```go
package main

import (
	"fmt"
	"github.com/nyudlts/go-aspace"
)

func main() {
    aspaceClient, err := aspace.NewClient(/path/to/go-aspace.yml, "environment from yaml", timeout)
    if err != nil {
        panic(err)
    }
    
    repositories, err := aspaceClient.GetRepositories()
    if err != nil {
        panic(err)
    }
	
    fmt.Println(repositories) //prints the array of Repository IDs from ArchivesSpace.
}
```