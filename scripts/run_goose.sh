#!/bin/bash

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgres://rssagg_user:password123@172.17.0.1:5432/rssagg?sslmode=disable
# shellcheck disable=SC2164
cd sql/schema
goose up