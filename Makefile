fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test -race ./...

build:
	go build ./cmd/app

all: fmt lint test build