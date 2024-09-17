#!/bin/bash

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgres://olehhuss:o2l0e0h4@localhost:5432/rssagg

goose up