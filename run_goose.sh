#!/bin/sh

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgres://postgres:root@localhost:5433/rssagg?sslmode=disable
# shellcheck disable=SC2164
cd sql/schema
goose up