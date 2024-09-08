
BUILD_DIR := build
BIN_NAME := golang-challenge

# Targets
.PHONY: all build clean

all: build

build:
	@mkdir -p $(BUILD_DIR)
	go build -o main ./cmd/$(BIN_NAME) 

build-lambda:
	@mkdir -p $(BUILD_DIR)
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o bootstrap main.go 

run:
	go run cmd/${BIN_NAME}/main.go

clean:
	@rm -rf $(BUILD_DIR)