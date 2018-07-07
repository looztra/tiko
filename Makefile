# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BUILD_DIR=build
BINARY_NAME=tiko
BINARY_LINUX=$(BINARY_NAME)-linux-x86_64

BIN_DIR:= $(shell echo $$GOPATH | cut -d ":" -f 1)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter

.PHONY: test

all: test build

build:
				CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_LINUX) -v

$(GOMETALINTER):
		go get -u github.com/alecthomas/gometalinter
		gometalinter --install &> /dev/null

.PHONY: lint
lint: $(GOMETALINTER)
		gometalinter ./... --vendor

test: lint
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -rf $(BUILD_DIR)

run: build
		./$(BINARY_UNIX)

deps:
		dep ensure
