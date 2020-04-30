# go-aspace
a Go library and utility for ArchivesSpace integrations

## Prerequisite
[GoLang](https://golang.org/dl/)

## Installation
1. Get the go-aspace library: $go get github.com/nyudlts/go-aspace.
2. Install the binary: $go install github.com/nyudlts/go-aspace.
3. Edit the config file `$GOROOT/github.com/nyudlts/go-aspace/go-aspace_template`, enter aspace credentials, and rename the file `go-aspace`
4. Move config file to `$HOME/`, `/etc/`, or `/etc/sysconfig/`, or leave it in `$GOROOT/src/github.com/nyudlts/go-aspace/`.

## Usage
  go-aspace [flags]<br>
  go-aspace [command]<br>

### Available Commands
#### export      
export a resource as EAD from an Archivesspace<br>
**Usage:**<br>
go-aspace export [flags]<br>
<br>
**Flags:**<br>
-d, --daos, bool              include daos (default true)<br>
-e, --ead3, bool               ead3 format (default false)<br>
-h, --help               help for export<br>
-l, --location, string    Location to write EAD File (default "/tmp")<br>
-n, --num_cs, bool             include numbered components (default false)<br>
-p, --pdf, bool                pdf format (default false)<br>
-r, --repositoryId, int   Id of the repository<br>
-c, --resourceId, int     Id of the resource (collection)<br>
-u, --unpub, bool              include unpublished (default false)<br>

