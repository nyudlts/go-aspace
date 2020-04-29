# go-aspace
a Go library and utility for ArchivesSpace integrations

## Prerequisite
GoLang

## Installation
1. Get the go-aspace library: $go get github.com/nyudlts/go-aspace.
2. Install the binary: $go install github.com/nyudlts/go-aspace.
3. Edit `go-aspace_template` and move to `$HOME/go-aspace`, `/etc/go-aspace`, or `/etc/sysconfig/go-aspace`

## Usage
  go-aspace [flags]<br>
  go-aspace [command]<br>

### Available Commands
  #### export      
  export a resource as EAD from archivesspace<br>
  **Usage:**<br>
    go-aspace export [flags]<br>
  <br>
  Flags:<br>
    **-h, --help**               help for export<br>
        **--location** string    Location to write EAD File (default "/tmp")<br>
        **--repositoryId** int   Id of the repository<br>
        **--resourceId int**     Id of the resource<br>
