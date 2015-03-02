## Golang application auto build versioning 

## Requirements 

You use tagging git capability to point importance of a release e.g. v1.0.0

## Usage

In some of your source code files e.g. main.go 

```golang
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
