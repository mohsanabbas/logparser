# Simple Makefile for managing Go projects

# Variables
BINARY_NAME=log-parser
BUILD_DIR=cmd/log-parser

# Set the GOBIN environment variable
export GOBIN=$(shell pwd)/bin

# Default target
all: build

# Build the binary without race detection
build:
	@echo "Building the binary..."
	go build -o bin/${BINARY_NAME} ${BUILD_DIR}/main.go

# Run the binary without race detection
run:
	@echo "Running the application..."
	go run ${BUILD_DIR}/main.go

# Run the binary with race detection
run-race:
	@echo "Running the application with race detector..."
	go run -race ${BUILD_DIR}/main.go

# Clean up the binary
clean:
	@echo "Cleaning up..."
	rm -f bin/${BINARY_NAME}

# Help
help:
	@echo "make        : Build the application."
	@echo "make run    : Run the application."
	@echo "make run-race : Run the application with the race detector enabled."
	@echo "make clean  : Remove binaries."
	@echo "make help   : Display this help."

.PHONY: build run run-race clean help
