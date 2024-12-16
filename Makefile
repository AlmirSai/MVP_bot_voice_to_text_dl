# Makefile

.PHONY: all build run clean test install lint

current_dir := $(shell pwd)

all: build

build:
	@echo "Building the project from $(current_dir)..."
	go build -o ./bot/main ./bot/cmd

run: build
	@echo "Running the project..."
	./bot/main

test:
	@echo "Running tests..."
	go test ./bot/...

clean:
	@echo "Cleaning the project..."
	rm -rf bot/main

install:
	@echo "Installing dependencies..."
	go mod tidy

lint:
	@echo "Running linter..."
	golangci-lint run
