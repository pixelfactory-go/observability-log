.PHONY: help fmt test lint build clean fuzz
SHELL := /bin/bash

## help: Display this help message
help:
	@echo "Available targets:"
	@echo "  fmt         - Format code using gofmt"
	@echo "  lint        - Run golangci-lint"
	@echo "  test        - Run tests with coverage"
	@echo "  fuzz        - Run all fuzz tests (or specific test with FUZZ_TEST=FuzzName)"
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
ifdef FUZZ_TEST
	@echo "Running fuzz test: $(FUZZ_TEST)..."
	@FUZZ_NAME=$$(echo "$(FUZZ_TEST)" | cut -d':' -f1); \
	FUZZ_PKG=$$(echo "$(FUZZ_TEST)" | cut -d':' -f2); \
	if [ -z "$$FUZZ_PKG" ] || [ "$$FUZZ_PKG" = "$$FUZZ_NAME" ]; then \
		FUZZ_PKG="."; \
	fi; \
	go test -run=^$$ -fuzz=$$FUZZ_NAME -fuzztime=30s $$FUZZ_PKG
	@echo "Done!"
else
	@echo "Running all fuzz tests..."
	@for file in $$(find . -name "*_fuzz_test.go" -o -name "*_test.go" | grep -E "(fuzz|_test\.go$$)"); do \
		pkg=$$(dirname $$file); \
		for test in $$(grep -h '^func Fuzz' $$file 2>/dev/null | sed 's/func \(Fuzz[^(]*\).*/\1/'); do \
			echo "Running $$test in $$pkg..."; \
			go test -run=^$$ -fuzz=$$test -fuzztime=30s $$pkg || exit 1; \
		done; \
	done
	@echo "All fuzz tests completed successfully!"
endif

## build: Build the project
build:
	@go build -v ./...

## clean: Clean build artifacts
clean:
	@rm -f coverage.txt
	@go clean -cache -testcache
