[![Build Status](https://travis-ci.com/NickMetz/golang-docker-sample.svg?branch=master)](https://travis-ci.com/NickMetz/)
[![Go Report Card](https://goreportcard.com/badge/github.com/NickMetz/golang-docker-sample)](https://goreportcard.com/report/github.com/NickMetz/golang-docker-sample)
[![Coverage Status](https://coveralls.io/repos/github/NickMetz/golang-docker-sample/badge.svg?branch=master)](https://coveralls.io/github/NickMetz/golang-docker-sample?branch=master)
[![Docker Pulls](https://img.shields.io/docker/pulls/nmetz/golang-docker-sample)](https://hub.docker.com/repository/docker/nmetz/golang-docker-sample)

# Golang Docker Sample

Golang Docker Sample is an example on how to test, build and publish a go app on Github.

## Quick Start

This repository produces a Docker image on Dockerhub:
```
# Start docker image
docker run nmetz/golang-docker-sample:latest
```

## Building and running

To build only sample app on your local machine it's required to have Golang Language installed. To install Golang please follow this instructions [here](https://golang.org/doc/install).
```
# Clone git repository
git clone https://github.com/NickMetz/golang-docker-sample
# Run go Build
go build cmd/golang-docker-sample/main.go
```
## Environment Variables

The following environment variables configure the Golang Docker Sample:

| Name | Description | Default |
|------|-------------|:-----:|
| `PARALLEL_JOBS` | Define how many jobs should run in parallel. | `5` |
| `LIMIT_JOBS` | Limit how many jobs should be executed in total. Default to `0` for infinity job execution. | `0` |