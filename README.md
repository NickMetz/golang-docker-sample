[![Build Status](https://travis-ci.com/NickMetz/golang-docker-sample.svg?branch=master)](https://travis-ci.com/NickMetz/)

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