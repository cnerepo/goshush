## goshush
Command line secrets.txt

## Usage

## Installation
A Go development environment is needed for developing and building this project.
Please download [Go v1.8.3](https://storage.googleapis.com/golang/go1.8.3.darwin-amd64.pkg) before proceeding with the Scurvy install.

Upon installing the Go runtime it will create a go directory under `$HOME/go`, which will be referred to as your $GOPATH.  This directory is where the Go runtime will expect your source code and any go related packages to live. If you wish to change this path to something that fits your development flow modify the GOPATH environment variable within your .profile or shell rc file(s) with the new path.

You can use the command `go env` to find your GOPATH and GOROOT paths.

```Bash
# Setting up GO envs and setting your PATH
mkdir -p $HOME/go/src

# Adding Go bin directories to your PATH in $HOME/.profile
export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export BASH=$PATH:$GOROOT/bin:$GOPATH/bin
```

### Install dependencies
```Bash
cd $GOPATH/src/cnerepo/goshush
go get -u ./...
```

### Project Cloning and Setup

```Bash
go build -o goshush cmd/main.go

```

