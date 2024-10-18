#!/usr/bin/env bash

set -eux

go install github.com/swaggo/swag/cmd/swag@latest
go install golang.org/x/tools/cmd/goimports@latest
