#!/usr/bin/env bash

set -eux

# swaggo is installed directly in dockerfile
# go install github.com/swaggo/swag/cmd/swag@latest

# mockery, gofumpt, and golangci-lint are installed with devcontainer features
# go install golang.org/x/tools/cmd/goimports@latest
# go install github.com/vektra/mockery/v2@v2.46.3
# curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.61.0
