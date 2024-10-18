.PHONY: build run tvet tidy fmt clean

build: vet
	go build

run: vet
	go run ./cmd/server

test: vet
	go test ./...

vet: tidy
	go vet ./...

tidy: fmt
	go mod tidy

fmt: swag
	go fmt ./...

swag:
	swag init --generalInfo ./cmd/server/main.go --output ./internal/docs
	swag fmt

clean:
	go clean

