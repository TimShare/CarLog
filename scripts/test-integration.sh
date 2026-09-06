#!/bin/sh

set -eu

project_name="carlog-tests"
compose_file="docker-compose.test.yml"

compose() {
    if command -v docker-compose >/dev/null 2>&1; then
        docker-compose "$@"
    else
        docker compose "$@"
    fi
}

cleanup() {
    compose -p "$project_name" -f "$compose_file" down --volumes --remove-orphans
}

trap cleanup EXIT INT TERM

cleanup
compose -p "$project_name" -f "$compose_file" up --build -d api

if [ "$#" -eq 0 ]; then
    compose -p "$project_name" -f "$compose_file" run --rm --build integration-tests
else
    compose -p "$project_name" -f "$compose_file" run --rm --build integration-tests "$@"
fi
