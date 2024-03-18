# Variables
BINARY_NAME=cephalopod
CMD_DIR=./cmd
BUILD_DIR=./bin

# Ensure the build directory exists
$(shell mkdir -p $(BUILD_DIR))

# Build the binary
build:
	@echo "Building..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

# Run the program
run: build
	@echo "Running..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	go clean
	rm -rf $(BUILD_DIR)

.PHONY: build run clean