CC=go

build:
	go build main.go

run:
	./main

run_with_args:
	./main -build=true
