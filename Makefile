.PHONY: build run test clean

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
BINARY_NAME=sd-gen-o2

# Build the main application
build: 
	$(GOBUILD) -o bin/$(BINARY_NAME) -v ./cmd/dummydata

# Run the application
run:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v ./cmd/dummydata
	./bin/$(BINARY_NAME)

# Run tests
test: 
	$(GOTEST) -v ./...

# Clean build artifacts
clean: 
	$(GOCLEAN)
	rm -rf bin/

# Download dependencies
deps:
	$(GOCMD) mod tidy
	$(GOCMD) mod download

# Build for cross-platform
build-all:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_NAME)-linux -v ./cmd/dummydata
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_NAME)-windows.exe -v ./cmd/dummydata
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_NAME)-mac -v ./cmd/dummydata