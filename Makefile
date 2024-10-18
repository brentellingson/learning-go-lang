.PHONY: build run tvet tidy fmt clean

build: lint
	go build

run: lint
	go run ./cmd/server

test: lint
	go test ./...

lint: fmt
	go vet ./...
	staticcheck ./...
	revive -formatter friendly ./...

fmt: swag
	go mod tidy
	go fmt ./...
	goimports -l -w .

swag:
	swag init --generalInfo ./cmd/server/main.go --output ./internal/docs
	swag fmt

clean:
	go clean

