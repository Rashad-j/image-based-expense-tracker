# Variables
APP_NAME := image-expense-tracker
APP_DIR := cmd/app
DOCKER_IMAGE := $(APP_NAME):latest
DOCKER_CONTAINER := $(APP_NAME)-container

# Default target: show help
.DEFAULT_GOAL := help

# Help documentation
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build          Build the Go application binary"
	@echo "  make run            Run the application locally"
	@echo "  make docker-build   Build a Docker image for the application"
	@echo "  make docker-run     Run the application in a Docker container"
	@echo "  make docker-stop    Stop the running Docker container"
	@echo "  make docker-clean   Remove Docker container and image"

# Build the Go binary
.PHONY: build
build: ## Build the Go application binary
	@echo "Building the Go application..."
	go build -o bin/$(APP_NAME) $(APP_DIR)/main.go

# Run the application locally
.PHONY: run
run: build ## Run the Go application locally
	@echo "Running the Go application locally..."
	./bin/$(APP_NAME)

# Build the Docker image
.PHONY: docker-build
docker-build: ## Build a Docker image for the application
	@echo "Building Docker image $(DOCKER_IMAGE)..."
	docker build -t $(DOCKER_IMAGE) .

# Run the Docker container
.PHONY: docker-run
docker-run: docker-build ## Run the application in a Docker container
	@echo "Running Docker container $(DOCKER_CONTAINER)..."
	docker run --rm -p 8080:8080 --name $(DOCKER_CONTAINER) $(DOCKER_IMAGE)

# Stop the Docker container
.PHONY: docker-stop
docker-stop: ## Stop the running Docker container
	@echo "Stopping Docker container $(DOCKER_CONTAINER)..."
	docker stop $(DOCKER_CONTAINER) || true

# Clean up Docker artifacts
.PHONY: docker-clean
docker-clean: ## Remove Docker container and image
	@echo "Removing Docker container $(DOCKER_CONTAINER)..."
	docker rm $(DOCKER_CONTAINER) || true
	@echo "Removing Docker image $(DOCKER_IMAGE)..."
	docker rmi $(DOCKER_IMAGE) || true

