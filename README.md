[![CircleCI](https://circleci.com/gh/spatialcurrent/go-stringify/tree/master.svg?style=svg)](https://circleci.com/gh/spatialcurrent/go-stringify/tree/master) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-stringify)](https://goreportcard.com/report/spatialcurrent/go-stringify)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/go-stringify?status.svg)](https://godoc.org/github.com/spatialcurrent/go-stringify) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/go-stringify/blob/master/LICENSE)

# go-stringify

# Description

**go-stringify** is a simple library for generating a stringer that converts literals to strings, converts map keys to strings, or converts a slice to a slice of strings.  **go-stringify** is used by [go-simple-serializer](http://github.com/spatialcurrent/go-simple-serializer).

# Usage

**Go**

You can install **go-stringify** with.


```shell
go get -u -d github.com/spatialcurrent/go-stringify/...
```

You can import the **stringify** package with:

```go
import (
  stringify "github.com/spatialcurrent/go-stringify/pkg/stringify"
)
```

See [stringify](https://godoc.org/github.com/spatialcurrent/go-stringify) in GoDoc for information on how to use Go API.

# Testing

Run test using `make test` or (`bash scripts/test.sh`), which runs unit tests, `go vet`, `go vet with shadow`, [errcheck](https://github.com/kisielk/errcheck), [ineffassign](https://github.com/gordonklaus/ineffassign), [staticcheck](https://staticcheck.io/), and [misspell](https://github.com/client9/misspell).

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-stringify/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
