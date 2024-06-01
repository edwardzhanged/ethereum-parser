.PHONY: build test clean

BINARY_NAME := ethereum-parser

# Build the project
build:
	go build -o $(BINARY_NAME) .

# Test the project
test:
	go test -v ./...

# Clean up
clean:
	go clean
	rm -f myapp