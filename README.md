
# GOGOLF Boilerplate Template
My go backend template. This is a go boiplerplate to quickly started the new project with golang and standard dependencies for local development pipeline.

**Features**

1. Hot-Reloader with Air.
2. Pre-commit.


# Getting Started
## Prerequsition
Please run `$ make init` to initialize the basic needed command and dependencies for your local development. We built this command to install the stuff to make the industry stnadard good grade to local development.

Here is a standard development pipeline installed.

1. go
2. node
3. pre-commit

## 👩‍⚕️ Pre-Commit
You can reached pre-commit hooks and rules under `.pre-commit-config.yaml` configuration file.

## Commit Lint
The commit lint is make more sense of good engineering should be. Essepcially, the `commitlint`, in this template I've follow the convention here [https://www.conventionalcommits.org/en/v1.0.0/](https://www.conventionalcommits.org/en/v1.0.0/)


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
