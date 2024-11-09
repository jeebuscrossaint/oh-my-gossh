# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install

# Binary name
BINARY_NAME=oh-my-gossh

# Build flags
LDFLAGS=-ldflags "-s -w"  # Strip debug info
GCFLAGS=-gcflags "-N -l"  # Disable optimizations for faster builds during dev
BUILDFLAGS=-trimpath     # Remove file system paths from binary

# Optimization flags for release builds
RELEASE_FLAGS=$(LDFLAGS) $(BUILDFLAGS) -tags release

.PHONY: all build clean install test help

# Default target
all: build

# Development build
build:
	CGO_ENABLED=0 $(GOBUILD) $(GCFLAGS) -o $(BINARY_NAME) .

# Optimized release build
release:
	CGO_ENABLED=0 $(GOBUILD) $(RELEASE_FLAGS) -o $(BINARY_NAME) .

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Install binary
install: release
	sudo cp $(BINARY_NAME) /usr/local/bin/

# Run tests
test:
	$(GOTEST) -v ./...

# Update dependencies
deps:
	$(GOGET) -u ./...

# Show help
help:
	@echo "Available targets:"
	@echo "  build    - Build development binary"
	@echo "  release  - Build optimized release binary"
	@echo "  clean    - Remove build artifacts"
	@echo "  install  - Install binary to system"
	@echo "  test     - Run tests"
	@echo "  deps     - Update dependencies"