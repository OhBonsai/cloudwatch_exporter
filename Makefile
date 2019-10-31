APP_NAME?=ace
PWD=$(shell pwd)
COMMIT_ID=$(shell git rev-parse HEAD)
BUILD_ID?=$(shell git rev-parse HEAD)
GOBUILD_FLAGS?=-i -ldflags "-X main.version=$(BUILD_ID)"

.DEFAULT_GOAL := ${APP_NAME}

.PHONY: ${APP_NAME}
${APP_NAME}:
	go build $(GOBUILD_FLAGS) -o $(APP_NAME) .

linux:
	CGO_ENABLED=0 GOOS=linux go build $(GOBUILD_FLAGS) -o $(APP_NAME) .



