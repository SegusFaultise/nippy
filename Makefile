# Define output binary and directories
BINARY_NAME=nippy
OUTPUT_DIR=bin
CMD_DIR=./cmd/nippy

# Define directories for testing
TEST_DIR=./test
PKG_DIR=./pkg

# Default target: build, test, and run the binary
.PHONY: all
all: build test run

# Build the Go binary
.PHONY: build
build:
	@echo "---------------------------------"
	@echo "🔨  Building the binary..."
	@echo "---------------------------------"
	go build -o $(OUTPUT_DIR)/$(BINARY_NAME) $(CMD_DIR)
	@echo "✅  Build complete!"
	@echo ""

# Run tests in the 'test' directory
.PHONY: test
test:
	@echo "---------------------------------"
	@echo "🧪  Running tests..."
	@echo "---------------------------------"
	go test -v $(TEST_DIR)/...
	@echo "✅  All tests complete!"
	@echo ""

# Run the binary if tests pass
.PHONY: run
run:
	@echo "---------------------------------"
	@echo "🚀  Running the binary..."
	@echo "---------------------------------"
	$(OUTPUT_DIR)/$(BINARY_NAME)
	@echo ""
	@echo "🏁  Program execution finished!"
	@echo ""

# Clean up binary files
.PHONY: clean
clean:
	@echo "---------------------------------"
	@echo "🧹  Cleaning up..."
	@echo "---------------------------------"
	rm -f $(OUTPUT_DIR)/$(BINARY_NAME)
	@echo "✅  Cleanup complete!"
	@echo ""

# Build, test, and run the binary in sequence
.PHONY: build-test-run
build-test-run: build test run
