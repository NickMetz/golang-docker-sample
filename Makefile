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
	go test ./cmd/golang-docker-sample/ -v -covermode=count -coverprofile=coverage.out
	$(HOME)/gopath/bin/goveralls -coverprofile=coverage.out -service=travis-ci

.PHONY: build
build:
	echo "$(DOCKER_PASSWORD)" | docker login -u "$(DOCKER_USERNAME)" --password-stdin
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

