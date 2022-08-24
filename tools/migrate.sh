#!/usr/bin/env bash

declare -r scriptPath="$(dirname "$0")"
declare -r scriptName="$(basename "$0")"

# show the usage description
if (( $# > 0 )); then
  echo "Usage:"
  echo "  $scriptName -h | --help"
  echo "  $scriptName"
  echo
  echo "Options:"
  echo "  -h, --help  - show the help message and exit."
  echo
  echo "Environment variables:"
  echo "  PATH_TO_MIGRATIONS  - path to migrations (default: \"$scriptPath/../migrations\");"
  echo "  STORAGE_ADDRESS     - PostgreSQL connection URI (default: \"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable\")."

  exit 0
fi

# declare and define the parameters
declare -r PATH_TO_MIGRATIONS_ON_HOST="${PATH_TO_MIGRATIONS:-"$scriptPath/../migrations"}"
declare -r PATH_TO_MIGRATIONS_IN_CONTAINER="/etc/go-exercises-backend/migrations"
declare -r STORAGE_ADDRESS="${STORAGE_ADDRESS:-"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"}"

# run applying migrations
declare -r absolutePathToMigrationsOnHost="$(realpath "$PATH_TO_MIGRATIONS_ON_HOST")"
docker run --volume "$absolutePathToMigrationsOnHost:$PATH_TO_MIGRATIONS_IN_CONTAINER" --network host \
  migrate/migrate:v4.15.2 -path="$PATH_TO_MIGRATIONS_IN_CONTAINER" -database "$STORAGE_ADDRESS" up
