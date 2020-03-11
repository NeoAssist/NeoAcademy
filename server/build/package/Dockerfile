FROM golang:1.13-alpine

COPY . /go/src/server
WORKDIR /go/src/server

RUN apk add --no-cache git && \
    go mod vendor
