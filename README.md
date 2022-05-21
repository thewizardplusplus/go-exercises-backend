# go-exercises-backend

[![GoDoc](https://godoc.org/github.com/thewizardplusplus/go-exercises-backend?status.svg)](https://godoc.org/github.com/thewizardplusplus/go-exercises-backend)
[![Go Report Card](https://goreportcard.com/badge/github.com/thewizardplusplus/go-exercises-backend)](https://goreportcard.com/report/github.com/thewizardplusplus/go-exercises-backend)
[![Build Status](https://app.travis-ci.com/thewizardplusplus/go-exercises-backend.svg?branch=master)](https://app.travis-ci.com/thewizardplusplus/go-exercises-backend)

Back-end of the service for solving programming exercises.

## Features

- RESTful API:
  - models:
    - task model:
      - storing:
        - author ID;
        - title;
        - description;
        - boilerplate code;
        - test cases:
          - all test cases are represented by a single string;
      - operations:
        - getting all tasks:
          - calculate a total correctness flag based on all solutions of a requesting author;
          - sort the results by creation time in descending order;
          - process pagination:
            - implemented using an offset and a limit;
        - getting a single task by an ID:
          - calculate a total correctness flag based on all solutions of a requesting author;
        - creating:
          - automatically format a task boilerplate code;
        - updating by an ID:
          - allowed for its author only;
          - automatically format a task boilerplate code;
        - deleting by an ID:
          - allowed for its author only;
    - solution model:
      - storing:
        - author ID;
        - task ID;
        - code;
        - correctness flag;
        - testing result:
          - represented by a string;
      - operations:
        - getting all solutions by a task ID:
          - filtered by a requesting author:
            - allow a solution task author to get solutions of other authors;
          - sort the results by creation time in descending order;
          - process pagination:
            - implemented using an offset and a limit;
        - getting a single solution by an ID:
          - allowed for:
            - solution author;
            - solution task author;
        - creating:
          - automatically format a solution code;
        - updating by an ID:
          - performed by a queue consumer only (see below);
        - formatting a solution code;
  - representing:
    - in a JSON:
      - payloads:
        - of requests;
        - of responses;
    - as a plain text:
      - errors;
- server:
  - additional routing:
    - serving static files;
  - storing settings in environment variables;
  - supporting graceful shutdown;
  - logging:
    - logging requests;
    - logging errors;
  - panics:
    - recovering on panics;
    - logging of panics;
- authentication:
  - use the Bearer authentication scheme based on [JSON Web Tokens](https://jwt.io/):
    - store in a JWT claims:
      - expiration time claim;
      - user claim:
        - contains a whole user model;
  - generate a token signing key automatically by default;
  - user model:
    - storing:
      - username (unique);
      - password hash:
        - generated using the [bcrypt](https://en.wikipedia.org/wiki/Bcrypt) function;
      - flag indicating whether the user is disabled or not;
- databases:
  - storing data in the [PostgreSQL](https://www.postgresql.org/) database;
- interaction with queues:
  - using the [RabbitMQ](https://www.rabbitmq.com/) message broker;
  - common properties:
    - automatic declaring of the used queues;
    - passing of a message data in JSON;
  - operations:
    - producing solutions:
      - concurrent producing;
    - consuming solution results:
      - concurrent consuming;
      - once requeue the solution on failure;
- utilities:
  - utility for managing users:
    - commands:
      - add a user:
        - parameters:
          - username;
          - password;
          - password hashing cost;
          - flag indicating the need to generate a password;
          - generated password length;
          - flag indicating whether the user is disabled or not;
      - update a user:
        - parameters:
          - username;
          - new username;
          - password;
          - password hashing cost;
          - flag indicating the need to generate a password;
          - generated password length;
          - flag indicating whether the user is disabled or not;
          - flag indicating whether the user is enabled or not;
        - can update the user fields individually;
- distributing:
  - [Docker](https://www.docker.com/) image;
  - [Docker Compose](https://docs.docker.com/compose/) configuration.

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

API description:

- RabbitMQ API description in the [AsyncAPI](https://www.asyncapi.com/) format: [docs/async_api.yaml](docs/async_api.yaml);
- web API description:
  - in the [Swagger](http://swagger.io/) format: [docs/swagger.yaml](docs/swagger.yaml);
  - in the format of a [Postman](https://www.postman.com/) collection: [docs/postman_collection.json](docs/postman_collection.json).

## Utilities

- [go-exercises-manager](cmd/go-exercises-manager) &mdash; utility for managing users

## License

The MIT License (MIT)

Copyright &copy; 2021-2022 thewizardplusplus
