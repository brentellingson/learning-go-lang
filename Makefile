.PHONY: build run tvet tidy fmt clean

build: lint
	go build

run: lint
	go run ./cmd/server

test: lint
	go test ./...

lint: format
	go vet ./...
	staticcheck ./...
	revive -formatter friendly ./...

format: generate
	go mod tidy
	go fmt ./...
	goimports -l -w .

generate:
	swag init --generalInfo ./cmd/server/main.go --output ./internal/docs
	swag fmt

clean:
	go clean

