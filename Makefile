
BUILD_DIR := build
BIN_NAME := golang-challenge

# Targets
.PHONY: all build clean

all: build

build:
	@mkdir -p $(BUILD_DIR)
	go build -o main ./cmd/$(BIN_NAME) 

run:
	go run cmd/${BIN_NAME}/main.go

clean:
	@rm -rf $(BUILD_DIR)