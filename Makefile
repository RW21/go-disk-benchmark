# Binary output directory
BINARY_DIR := ./bin
# Binary output name
BINARY_NAME := myapp

# Go files
GO_FILES := $(shell find . -type f -name '*.go')

# Go build flags
GO_BUILD_FLAGS := -v

# Default make command
default: build

# Build the binary
build: $(GO_FILES)
	@echo "  >  Building binary..."
	go build $(GO_BUILD_FLAGS) -o $(BINARY_DIR)/$(BINARY_NAME) ./cmd/main/main.go

# Clean the binary
clean:
	@echo "  >  Cleaning build cache"
	go clean
	rm -rf $(BINARY_DIR)

.PHONY: default build clean
