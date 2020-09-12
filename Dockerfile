FROM golang:1.14-alpine3.12 AS build

WORKDIR /build

COPY ./ .
RUN go build cmd/golang-docker-sample/main.go

FROM alpine:latest

WORKDIR /app

RUN addgroup -S appuser && adduser -S appuser -G appuser
COPY --chown=appuser:appuser --from=build /build/main /app/golang-docker-sample

USER appuser
ENTRYPOINT [ "/app/golang-docker-sample" ]