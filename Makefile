.PHONY: build run clean run-tests

# Variables
BINARY_NAME=squid
SOURCE_DIR=./
BUILD_DIR=./bin

# Ensure the build directory exists
$(shell mkdir -p $(BUILD_DIR))

# Build the binary
build:
	rm -rf ./bin
	@echo "Building..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SOURCE_DIR)

# Run the program
run: build
	@echo "Running..."
	$(BUILD_DIR)/$(BINARY_NAME)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	go clean
	rm -rf $(BUILD_DIR)

run-tests:
	@echo "Running tests..."
	go test -v ./...