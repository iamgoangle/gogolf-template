
# GOGOLF Boilerplate Template
My go backend template. This is a go boiplerplate to quickly started the new project with golang and standard dependencies for local development pipeline.

**Features**

1. Hot-Reloader with Air.
2. Pre-commit.


# Getting Started
## Prerequsition
Please run `$ make init` to initialize the basic needed command and dependencies for your local development. We built this command to install the stuff to make the industry standard good grade to local development.

Here is a standard development pipeline installed.

1. go
2. node
3. pre-commit

## Pre-Commit
1. commit linter
2. code formatter
3. unit test
4. code scanner
5. hooks

## Miscellaneous
### Set $GOPATH
The $GOPATH environment variable lists places for Go to look for Go Workspaces.

[https://go.dev/doc/gopath_code#GOPATH](https://go.dev/doc/gopath_code#GOPATH)

Add this to your profile. If you use *zsh* please also update `.zshrc`

```sh
export GOPATH="$HOME/go"
PATH="$GOPATH/bin:$PATH"
```

then `source .zshrc` for reload the the environment. You can check you GOPATH by `go env GOPATH`