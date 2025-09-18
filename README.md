# go-aspace v0.8.0
a Go library for ArchivesSpace integrations

## Use
1. Get the go-aspace library: $go get github.com/nyudlts/go-aspace.
2. Edit the config file `$GOROOT/github.com/nyudlts/go-aspace/go-aspace.yml_template`, enter your aspace credentials, and save it somewhere as `go-aspace.yml`
4. Import "github.com/nyudlts/go-aspace" into your project
5. Create an instance of the aspace client:

```go
package main

import (
	"fmt"
	"github.com/nyudlts/go-aspace"
)

func main() {
    aspaceClient, err := aspace.NewClient(/path/to/go-aspace.yml, "environment to use from config")
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
  timeout: 20
  
local:
  url: http://localhost:8089
  username: your-username
  password: your-password
  timeout: 20
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
**!!! Note** the test suite is designed to be run against a blank archivesspace database, it will create and delete repository, resource, archival_object, etc. do not use on an aspace database containing any actual data it will blank all repositories and resopurces prior to starting tests
