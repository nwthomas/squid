# Makefile for squid
.PHONY: build run clean test

# Variables
BINARY_NAME=squid
SOURCE_DIR=./
BUILD_DIR=./bin

# Ensure the build directory exists
$(shell mkdir -p $(BUILD_DIR))

# Build the binary
build:
	@echo "Building..."
	rm -rf $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SOURCE_DIR)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	go clean
	@rm -rf $(BUILD_DIR)

# Run the program
run: build
	@echo "Running..."
	$(BUILD_DIR)/$(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...
