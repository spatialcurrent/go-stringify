[![CircleCI](https://circleci.com/gh/spatialcurrent/go-stringify/tree/main.svg?style=svg)](https://circleci.com/gh/spatialcurrent/go-stringify/tree/main)
[![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-stringify?style=flat-square)](https://goreportcard.com/report/github.com/spatialcurrent/go-stringify)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/spatialcurrent/go-stringify)](https://pkg.go.dev/github.com/spatialcurrent/go-stringify)
[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/go-stringify/blob/master/LICENSE)

# go-stringify

# Description

**go-stringify** is a simple library for generating a stringer that converts literals to strings, converts map keys to strings, or converts a slice to a slice of strings.  **go-stringify** is used by [go-simple-serializer](http://github.com/spatialcurrent/go-simple-serializer).

# Usage

**Go**

You can import **go-stringify** as a library with:

```go
import (
  "github.com/spatialcurrent/go-stringify/pkg/stringify"
)
```

See [stringify](https://pkg.go.dev/github.com/spatialcurrent/go-stringify/pkg/stringify) in the docs for information on how to use Go API.  See the tests for ways to use this library.

# Testing

To run Go tests using `make test_go` or (`bash scripts/test.sh`), which runs unit tests, `go vet`, `go vet with shadow`, [errcheck](https://github.com/kisielk/errcheck), [staticcheck](https://staticcheck.io/), and [misspell](https://github.com/client9/misspell).

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-try-get/blob/main/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
