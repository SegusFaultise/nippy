CC=go

build:
	go build -o bin/nippy ./cmd/nippy

run:
	go run ./cmd/nippy

test:
	go test ./..
