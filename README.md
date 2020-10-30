# go-aspace
a Go library for ArchivesSpace integrations

## Use
1. Get the go-aspace library: $go get github.com/nyudlts/go-aspace.
2. Edit the config file `$GOROOT/github.com/nyudlts/go-aspace/go-aspace_template`, enter aspace credentials
3. Move config file to `$HOME/go-aspace`, `/etc/go-aspace`, or `/etc/sysconfig/go-aspace`, or leave it in `$GOROOT/src/github.com/nyudlts/go-aspace/go-aspace`.
4. Import "github.com/nyudlts/go-aspace"
5. An instance of the aspace client is now accessible at aspace.Client

```go
package main

import (
	"fmt"
	"github.com/nyudlts/go-aspace"
)

func main() {
	aspaceClient := aspace.Client
	repositories, err := aspaceClient.GetRepositoryList()
	if err != nil {
		panic(err)
	}
	fmt.Println(repositories) //prints the array of Repository IDs from ArchivesSpace.
}
```