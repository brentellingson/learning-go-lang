.PHONY: docker-run docker-build build run test lint format generate clean

docker-run: docker-build
	docker run -p 8080:8080 learning-golang-api

docker-build: format
	docker build -t learning-golang-api .

build: format
	go build

run: format
	go run ./cmd/server

test: format
	go test ./...  -cover

lint: format
	golangci-lint run

format: generate
	swag fmt
	gofumpt -l -w .

generate: restore
	swag init --generalInfo ./cmd/server/main.go --output ./internal/docs
	mockery

restore:
	go get -v ./...
	go mod tidy -v

