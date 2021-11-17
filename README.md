# go-exercises-backend

[![GoDoc](https://godoc.org/github.com/thewizardplusplus/go-exercises-backend?status.svg)](https://godoc.org/github.com/thewizardplusplus/go-exercises-backend)
[![Go Report Card](https://goreportcard.com/badge/github.com/thewizardplusplus/go-exercises-backend)](https://goreportcard.com/report/github.com/thewizardplusplus/go-exercises-backend)
[![Build Status](https://app.travis-ci.com/thewizardplusplus/go-exercises-backend.svg?branch=master)](https://app.travis-ci.com/thewizardplusplus/go-exercises-backend)

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

- `SERVER_STATIC_FILE_PATH` &mdash; path to static files (default: `./static`);
- addresses:
  - `SERVER_ADDRESS` &mdash; server URI (default: `:8080`);
  - `STORAGE_ADDRESS` &mdash; [PostgreSQL](https://www.postgresql.org/) connection URI (default: `postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable`);
  - `MESSAGE_BROKER_ADDRESS` &mdash; [RabbitMQ](https://www.rabbitmq.com/) connection URI (default: `amqp://rabbitmq:rabbitmq@localhost:5672`);
- solution registration:
  - `SOLUTION_REGISTER_BUFFER_SIZE` &mdash; solution registration channel capacity (default: `1000`);
  - `SOLUTION_REGISTER_CONCURRENCY` &mdash; amount of solution registration threads (default: `1000`);
- authorization:
  - `AUTHORIZATION_TOKEN_SIGNING_KEY` &mdash; authorization token signing key (is generated automatically if empty; default: empty);
  - `AUTHORIZATION_TOKEN_TTL` &mdash; authorization token TTL (default: `24h`).

## API Description

API description in the format of a [Postman](https://www.postman.com/) collection: [docs/postman_collection.json](docs/postman_collection.json).

## License

The MIT License (MIT)

Copyright &copy; 2021 thewizardplusplus
