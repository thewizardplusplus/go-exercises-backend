# go-exercises-backend

[![GoDoc](https://godoc.org/github.com/thewizardplusplus/go-exercises-backend?status.svg)](https://godoc.org/github.com/thewizardplusplus/go-exercises-backend)
[![Go Report Card](https://goreportcard.com/badge/github.com/thewizardplusplus/go-exercises-backend)](https://goreportcard.com/report/github.com/thewizardplusplus/go-exercises-backend)
[![Build Status](https://travis-ci.org/thewizardplusplus/go-exercises-backend.svg?branch=master)](https://travis-ci.org/thewizardplusplus/go-exercises-backend)

## Installation

Prepare the directory:

```
$ mkdir --parents "$(go env GOPATH)/src/github.com/thewizardplusplus/"
$ cd "$(go env GOPATH)/src/github.com/thewizardplusplus/"
```

Clone this repository:

```
$ git clone https://github.com/thewizardplusplus/go-exercises-backend.git
$ cd go-exercises-backend
```

Install dependencies with the [dep](https://golang.github.io/dep/) tool:

```
$ dep ensure -vendor-only
```

Build the project:

```
$ go install ./...
```

## Usage

```
$ go-exercises-backend
```

Environment variables:

- `SERVER_ADDRESS` &mdash; server URI (default: `:8080`);
- `STORAGE_ADDRESS` &mdash; [PostgreSQL](https://www.postgresql.org/) connection URI (default: `postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable`).

## API Description

API description in the format of a [Postman](https://www.postman.com/) collection: [docs/postman_collection.json](docs/postman_collection.json).

## License

The MIT License (MIT)

Copyright &copy; 2021 thewizardplusplus
