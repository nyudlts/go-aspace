# go-aspace
a Go library for ArchivesSpace integrations

## Use
1. Get the go-aspace library: $go get github.com/nyudlts/go-aspace.
2. Edit the config file `$GOROOT/github.com/nyudlts/go-aspace/go-aspace.yml_template`, enter aspace credentials
3. Move config file to `/etc/go-aspace.yml`.
4. Import "github.com/nyudlts/go-aspace"
5. Create an instance of the aspace client:

```go
package main

import (
	"fmt"
	"github.com/nyudlts/go-aspace"
)

func main() {
    aspaceClient, err := aspace.NewClient("environment", timeout)
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