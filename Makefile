.PHONY: help fmt test lint build clean fuzz fuzz-single
SHELL := /bin/bash

## help: Display this help message
help:
	@echo "Available targets:"
	@echo "  fmt         - Format code using gofmt"
	@echo "  lint        - Run golangci-lint"
	@echo "  test        - Run tests with coverage"
	@echo "  fuzz        - Run all fuzz tests (30s each)"
	@echo "  fuzz-single - Run a single fuzz test (FUZZ_TEST=FuzzName:package)"
	@echo "  build       - Build the project"
	@echo "  clean       - Clean build artifacts"
	@echo "  help        - Display this help message"

## fmt: Format code using gofmt
fmt:
	@diff -u <(echo -n) <(gofmt -d -s .)

## lint: Run golangci-lint
lint:
	@golangci-lint run ./...

## test: Run tests with coverage
test:
	@go test -v -race -coverprofile coverage.txt -covermode atomic ./...

## fuzz: Run fuzz tests for 30 seconds each
fuzz:
	@echo "Running fuzz tests (30s each)..."
	@go test -fuzz=FuzzLogMessages -fuzztime=30s .
	@go test -fuzz=FuzzLogLevel -fuzztime=30s .
	@go test -fuzz=FuzzLoggerWithFields -fuzztime=30s .
	@go test -fuzz=FuzzNewLogger -fuzztime=30s .
	@go test -fuzz=FuzzLoggerMultipleFields -fuzztime=30s .
	@go test -fuzz=FuzzUserAgent -fuzztime=30s ./fields
	@go test -fuzz=FuzzService -fuzztime=30s ./fields
	@go test -fuzz=FuzzSource -fuzztime=30s ./fields
	@go test -fuzz=FuzzURL -fuzztime=30s ./fields
	@go test -fuzz=FuzzHTTPRequest -fuzztime=30s ./fields
	@go test -fuzz=FuzzCoreWrite -fuzztime=30s ./sentry
	@go test -fuzz=FuzzCoreWriteWithFields -fuzztime=30s ./sentry
	@go test -fuzz=FuzzCoreWriteWithError -fuzztime=30s ./sentry
	@echo "All fuzz tests completed successfully!"

## fuzz-single: Run a single fuzz test (usage: make fuzz-single FUZZ_TEST=FuzzName:package)
fuzz-single:
	@if [ -z "$(FUZZ_TEST)" ]; then \
		echo "Error: FUZZ_TEST is required. Usage: make fuzz-single FUZZ_TEST=FuzzName:package"; \
		exit 1; \
	fi
	@FUZZ_NAME=$$(echo "$(FUZZ_TEST)" | cut -d':' -f1); \
	FUZZ_PKG=$$(echo "$(FUZZ_TEST)" | cut -d':' -f2); \
	echo "Running fuzz test: $$FUZZ_NAME in package $$FUZZ_PKG"; \
	go test -fuzz="^$$FUZZ_NAME$$" -fuzztime=30s $$FUZZ_PKG

## build: Build the project
build:
	@go build -v ./...

## clean: Clean build artifacts
clean:
	@rm -f coverage.txt
	@go clean -cache -testcache
