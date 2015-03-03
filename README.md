## Golang application auto build versioning 

## Requirements 

Package usage requires you to use tagging git capability to point importance of a release.
Make sure to use semantic versioning when you create a new tag.
A tag may include *v* prefix e.g. v1.0.0-alpha. 

When no tag is found, and you still use git, it creates a version 0.0.0-prerelease.  

## Usage

In a source code e.g. *main.go* file 

```go
package main

import (
  "fmt"

  "github.com/iaintshine/go-version/version"
)

var (
  build
  buildVersion Version = version.ParseGitDescription(build)
)  

func main() {
  fmt.Printf("Build info: %v", buildVersion)
}
```

And the Makefile

```
VERSION := $(shell git describe --tags)

install:
    go install -ldflags "-X github.com/<rep-path>/main.build ${VERSION}" github.com/<repo-path>
```
