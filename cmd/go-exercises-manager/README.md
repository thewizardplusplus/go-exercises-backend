# go-exercises-manager

The utility for managing users.

## Features

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
    - can update the user fields individually.

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

Build the utility:

```
$ go install ./cmd/go-exercises-manager
```

## Usage

```
$ go-exercises-manager -h | --help
$ go-exercises-manager add-user [options]
$ go-exercises-manager update-user [options]
```

Commands:

- `add-user` &mdash; add the user;
- `update-user` &mdash; update the user.

Options:

- `-h`, `--help` &mdash; show the context-sensitive help message and exit;
- `-u STRING`, `--username STRING` &mdash; username;
- `-U STRING`, `--new-username STRING` &mdash; new username (applies only to the `update-user` command);
- `-p STRING`, `--password STRING` &mdash; user password;
- `-c INTEGER`, `--cost INTEGER` &mdash; cost of the user password hashing (range: from 4 to 31 inclusive; default: `10`);
- `-g`, `--generate` &mdash; generate the user password;
- `-l INTEGER`, `--length INTEGER` &mdash; length of the user password to be generated (minimum: 6; default: `6`);
- `-d`, `--disable` &mdash; disable the user;
- `-e`, `--enable` &mdash; enable the user (applies only to the `update-user` command).

Environment variables:

- `STORAGE_ADDRESS` &mdash; [PostgreSQL](https://www.postgresql.org/) connection URI (default: `postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable`).
