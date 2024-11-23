# Variables
BINARY_NAME = ocr_parser
SRC_DIR = ./cmd/app/main.go
BUILD_DIR = ./bin

# Default target
all: build run

# Build the binary
build:
	@echo "Building the binary..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)/main.go
	@echo "Build complete. Binary located at $(BUILD_DIR)/$(BINARY_NAME)"

# Run the binary
run: build
	@echo "Running the binary..."
	@$(BUILD_DIR)/$(BINARY_NAME)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete."

# Help message
help:
	@echo "Available commands:"
	@echo "  make build    - Build the binary"
	@echo "  make run      - Build and run the binary"
	@echo "  make clean    - Remove build artifacts"
	@echo "  make help     - Show this help message"

.PHONY: all build run clean help
