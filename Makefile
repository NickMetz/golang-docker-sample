VERSION := 0.1.0
APP_NAME := golang-docker-sample

# Docker image name for this project
REPOSITORY := nmetz/$(APP_NAME)

# Shell to use for running scripts
SHELL := $(shell which bash)

# Commit hash from git
COMMIT=$(shell git rev-parse HEAD)

.PHONY: test
test:
	go test -v ./cmd/golang-docker-sample/

.PHONY: build
build:
	docker build \
	    -t $(APP_NAME) \
		-t $(REPOSITORY):latest \
		-t $(REPOSITORY):$(COMMIT) \
		-t $(REPOSITORY):$(VERSION) \
		-f Dockerfile \
		.

.PHONY: push
push:
	docker push $(REPOSITORY):$(VERSION)

.PHONY: release
release: test build push
