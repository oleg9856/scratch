#!/bin/sh

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgres://root:root@postgres_container:5432/rssagg?sslmode=disable
# shellcheck disable=SC2164
cd sql/schema
goose up