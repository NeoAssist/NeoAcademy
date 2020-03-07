FROM golang:1.13-alpine

COPY . /go/src/server/

WORKDIR /go/src/server/cmd/server

RUN apk add --no-cache git

RUN go get -d -v ./...
RUN go install -v ./...