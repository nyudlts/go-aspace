# go-aspace
a Go library and utility for ArchivesSpace integrations

## Binary Installation
1. wget https://github.com/nyudlts/go-aspace/releases/download/v0.2.0/go-aspace-linux or https://github.com/nyudlts/go-aspace/releases/download/v0.2.0/go-aspace-mac
2. sudo mv go-aspace-* /usr/local/bin
3. sudo chmod +x /usr/local/bin/go-aspace
4. wget https://github.com/nyudlts/go-aspace/releases/download/v0.2.0/go-aspace_template
5. mv go-aspace_template $HOME/go-aspace 
6. edit $HOME/go-aspace enter your archivesspace credentials
 
## Source Installation
1. Get the go-aspace library: $go get github.com/nyudlts/go-aspace.
2. Install the binary: $go install github.com/nyudlts/go-aspace.
3. Edit the config file `$GOROOT/github.com/nyudlts/go-aspace/go-aspace_template`, enter aspace credentials, and rename the file `go-aspace`
4. Move config file to `$HOME/`, `/etc/`, or `/etc/sysconfig/`, or leave it in `$GOROOT/src/github.com/nyudlts/go-aspace/`.

## Usage
  go-aspace [flags]<br>
  go-aspace [command]<br>

### Available Commands
* export
* sample
* version
* validate

#### export      
Export a resource as EAD from an Archivesspace<br>

**Usage:**<br>
go-aspace export [flags]<br>

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

#### sample     
Generate a sample set of resources as EAD from ArchivesSpace<br>

**Usage:**<br>
go-aspace sample [flags]<br>

**Flags:**<br>
  -h, --help                  help for sample<br>
  -l, --location string       Location to write EAD Files (default "/tmp")<br>
  -r, --repositories string   List of repository ids to be included in sample set (default "2")<br>
  -s, --size int              Size of the sample (default 1)<br>
  
#### validate
Validate all resources in a repository

**Usage:**<br>
go-aspace validate [flags]<br>


**Flags**:<br>
  -f, --filename string       Name of output file (default "go-aspace-validator.tsv")<br>
  -h, --help                  help for validate<br>
  -l, --location string       Location to write EAD Files (default "/tmp")<br>
  -p, --published             Validate only published resources (default true)<br>
  -r, --repositories string   List of repository ids to be included in sample set (default "2")<br>


