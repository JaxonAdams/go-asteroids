# Binary output name
BINARY_NAME=my-app

# Go source files
GO_FILES=$(wildcard *.go)

all: deps build

deps:
	go mod tidy

build:
	go build -o bin/$(BINARY_NAME).exe $(GO_FILES)

run: build
	./bin/$(BINARY_NAME).exe

clean:
	rm -rf bin

.PHONY: all deps build clean
.DEFAULT_GOAL := build
