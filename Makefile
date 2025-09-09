
BINARY_NAME ?= todo

.PHONY: all build test clean run deps build-linux

all: test build

build:
	mkdir -p bin
	go build -o bin/$(BINARY_NAME) cmd/todo/main.go

clean:
	rm -f bin/$(BINARY_NAME)
	rm -f data/*

run:
	go build -o bin/$(BINARY_NAME) cmd/todo/main.go
	./bin/$(BINARY_NAME)
