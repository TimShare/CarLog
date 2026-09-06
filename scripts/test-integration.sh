#!/bin/sh

set -eu

project_name="carlog-tests"
compose_file="docker-compose.test.yml"

cleanup() {
    docker-compose -p "$project_name" -f "$compose_file" down --volumes --remove-orphans
}

trap cleanup EXIT INT TERM

cleanup
docker-compose -p "$project_name" -f "$compose_file" up --build -d api

if [ "$#" -eq 0 ]; then
    docker-compose -p "$project_name" -f "$compose_file" run --rm --build integration-tests
else
    docker-compose -p "$project_name" -f "$compose_file" run --rm --build integration-tests "$@"
fi
