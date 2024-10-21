#!/usr/bin/env bash

set -eux

go install github.com/swaggo/swag/cmd/swag@latest
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/vektra/mockery/v2@v2.46.3
