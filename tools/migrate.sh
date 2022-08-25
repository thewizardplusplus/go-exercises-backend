#!/usr/bin/env bash

declare -r scriptPath="$(dirname "$0")"

# declare and define the default parameter values
declare -r DEFAULT_PATH_TO_MIGRATIONS="$scriptPath/../migrations"
declare -r DEFAULT_STORAGE_ADDRESS="postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

# show the usage description
declare -r normalizedDefaultPathToMigrations="./$(realpath --relative-to . "$DEFAULT_PATH_TO_MIGRATIONS")"
if (( $# > 0 )); then
  echo "Usage:"
  echo "  $0 -h | --help"
  echo "  $0"
  echo
  echo "Options:"
  echo "  -h, --help  - show the help message and exit."
  echo
  echo "Environment variables:"
  echo "  PATH_TO_MIGRATIONS  - path to migrations (default: \"$normalizedDefaultPathToMigrations\");"
  echo "  STORAGE_ADDRESS     - PostgreSQL connection URI (default: \"$DEFAULT_STORAGE_ADDRESS\")."

  exit 0
fi

# declare and define the parameters
declare -r PATH_TO_MIGRATIONS_ON_HOST="${PATH_TO_MIGRATIONS:-"$normalizedDefaultPathToMigrations"}"
declare -r PATH_TO_MIGRATIONS_IN_CONTAINER="/etc/go-exercises-backend/migrations"
declare -r STORAGE_ADDRESS="${STORAGE_ADDRESS:-"$DEFAULT_STORAGE_ADDRESS"}"

# run applying migrations
declare -r absolutePathToMigrationsOnHost="$(realpath "$PATH_TO_MIGRATIONS_ON_HOST")"
docker run --volume "$absolutePathToMigrationsOnHost:$PATH_TO_MIGRATIONS_IN_CONTAINER" --network host \
  migrate/migrate:v4.15.2 -path="$PATH_TO_MIGRATIONS_IN_CONTAINER" -database "$STORAGE_ADDRESS" up
