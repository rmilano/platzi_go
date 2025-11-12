GO ?= go
APP_NAME := HolaMundo
BIN_DIR := bin
SRC_DIR := HolaMundo
BINARY := $(BIN_DIR)/$(APP_NAME)

.PHONY: build run test clean fmt vet

build:
	mkdir -p $(BIN_DIR)
	$(GO) build -o $(BINARY) ./$(SRC_DIR)

run:
	$(GO) run ./$(SRC_DIR)

test:
	$(GO) test ./...

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

clean:
	rm -rf $(BIN_DIR)
