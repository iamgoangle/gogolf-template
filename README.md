# GOGOLF Boilerplate Template
My go backend template. This is a go boiplerplate to quickly started the new project with golang and standard dependencies for local development pipeline.
 ![Tux, the Linux mascot](/docs/GOGolf.jpg)
 
<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

**Table of Contents**
- [Getting Started](#getting-started)
- [Prerequsition](#prerequsition)
- [👩‍⚕️ Pre-Commit](#%E2%80%8D-pre-commit)
- [Commit Lint](#commit-lint)
- [Miscellaneous](#miscellaneous)
- [Set $GOPATH](#set-gopath)
- [Hexagonal Architecture](#hexagonal-architecture)
- [Architecture Components](#architecture-components)
- [Core](#core)
- [Domain](#domain)
- [Service](#service)
- [Repository](#repository)
- [Dependency Injection](#dependency-injection)
  - [Port and Adaptor Design Patterns](#port-and-adaptor-design-patterns)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

**Design Structure**
[สรุปเรื่อง GopherCon 2018: Kat Zien — How Do You Structure Your Go Apps](https://goangle.medium.com/%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87-gophercon-2018-kat-zien-how-do-you-structure-your-go-apps-a96faca9e8f1)

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

# Hexagonal Architecture
This golang project adopt the hexagonal architecture to grouping the code by domain and context of business problems.

## Architecture Components
### Core

Everything surrounded the `core` this layer contains the concrete core business logics.

The core could be viewed as a “box” (represented as a hexagon) capable of resolve all the business logic independently.

### Domain

The domain entity it contains the go struct definition of each entity that is part of the domain problem and can be used across the application.

>**Seperate of concern**
For example, this layer should not be known how the data keeps its, or how the caching store in the redis.

**Port** is the interfaces are the ports that external actors will plug their adapters into to drive the application.

**Adapter** is the implementation of the ports

### Service
Service is through which the client will interact with actual business logic/domain objects. The client can be rest handlers, test agents, etc.

### Repository
The repository interface allows us to connect to multiple data sources.

### Dependency Injection
#### Port and Adaptor Design Patterns

Outside concerns, such as repositories and user interfaces, use the adapters to plug into the business domain services.

**References**
- https://www.linkedin.com/pulse/part-2-go-project-layout-hexagonal-architecture-kaushal-prajapati/
- https://www.linkedin.com/pulse/hexagonal-software-architecture-implementation-using-golang-ramaboli/
- https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3
