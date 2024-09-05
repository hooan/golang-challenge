
GO := go
BUILD_DIR := build
BIN_NAME := golang-challenge

# Targets
.PHONY: all build clean

all: build

build:
	@mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(BIN_NAME) .

run:
	$(GO) run cmd/${BIN_NAME}/main.go

clean:
	@rm -rf $(BUILD_DIR)