.PHONY: all build run clean test install lint docker docker-run

current_dir := $(shell pwd)

# Default target
all: build

# Build the project
build:
	@echo "Building the project from $(current_dir)..."
	go build -o ./bot/main ./bot/cmd

# Run the project
run: build
	@echo "Running the project..."
	./bot/main

# Run tests
test:
	@echo "Running tests..."
	go test ./bot/...

# Clean build artifacts
clean:
	@echo "Cleaning the project..."
	rm -rf bot/main

# Install dependencies
install:
	@echo "Installing dependencies..."
	go mod tidy

# Lint the project
lint:
	@echo "Running linter..."
	golangci-lint run

# Build the Docker image
docker:
	@echo "Building Docker image..."
	docker build -t bot-image .

# Run the Docker container
docker-run: docker
	@echo "Running Docker container..."
	docker run --rm -it bot-image
