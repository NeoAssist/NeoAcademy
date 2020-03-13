FROM golang:1.13-alpine

COPY . /go/src/server/

WORKDIR /go/src/server/cmd/neoacademy

RUN apk add --no-cache git && \
    go mod vendor