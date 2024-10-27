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
	mockery

clean:
	go clean

