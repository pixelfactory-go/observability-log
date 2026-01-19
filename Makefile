.PHONY: help fmt test lint build clean
SHELL := /bin/bash

## help: Display this help message
help:
	@echo "Available targets:"
	@echo "  fmt     - Format code using gofmt"
	@echo "  lint    - Run golangci-lint"
	@echo "  test    - Run tests with coverage"
	@echo "  build   - Build the project"
	@echo "  clean   - Clean build artifacts"
	@echo "  help    - Display this help message"

## fmt: Format code using gofmt
fmt:
	@diff -u <(echo -n) <(gofmt -d -s .)

## lint: Run golangci-lint
lint:
	@golangci-lint run ./...

## test: Run tests with coverage
test:
	@go test -v -race -coverprofile coverage.txt -covermode atomic ./...

## build: Build the project
build:
	@go build -v ./...

## clean: Clean build artifacts
clean:
	@rm -f coverage.txt
	@go clean -cache -testcache
