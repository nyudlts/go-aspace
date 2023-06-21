# go-aspace v0.3.14b
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
    aspaceClient, err := aspace.NewClient(/path/to/go-aspace.yml, "environment to use from config", timeout)
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

## Config Format
A config template can be found in go-aspace.yml_template

```
dev:
  url: https://your-archivesspace-api.url:8089
  username: your-username
  password: your-password
  
local:
  url: https:/localhost:8089
  username: your-username
  password: your-password
```

## Example
A simple example application can be found at /example/main.go

```shell
go run example/main.go --config /path/to/go-aspace.yml --environment the-environment-to-use
```

## Testing
to run the test suite
```shell
go test -v --config /path/to/go-aspace.yml --environment the-environment-to-use
```
The test suite will select random objects from the specified aspace instance, which may fail serialization to go structs