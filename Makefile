# Binary name
BINARY_NAME=app

# Go related variables
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

# Ensure GOPATH is set before running build process
GOPATH?=$(shell go env GOPATH)

# Add .PHONY declaration for all targets that don't represent files
.PHONY: all build clean test run fmt help

# Default target
all: build

# Format the code
fmt:
	@echo "Running go fmt..."
	@go fmt ./...

# Build the binary
build: fmt
	@echo "Building..."
	@go build -o $(GOBIN)/$(BINARY_NAME) .

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(GOBIN)
	@go clean

# Run the binary
run: build
	@echo "Running binary..."
	@$(GOBIN)/$(BINARY_NAME)

# Help target
help:
	@echo "Available targets:"
	@echo "  all    - Default target, same as build"
	@echo "  fmt    - Format code using go fmt"
	@echo "  build  - Format code and build binary"
	@echo "  test   - Run tests"
	@echo "  clean  - Remove built artifacts"
	@echo "  run    - Build and run the binary"
	@echo "  help   - Show this help message"
