#!/usr/bin/env bash
set -e

GOOS=linux CGO_ENABLED=0 GOGC=off  go build -v -a -installsuffix nocgo -o dist/app ./src