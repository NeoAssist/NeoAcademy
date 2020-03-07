FROM node:13-alpine

RUN mkdir -p /usr/src/client
WORKDIR /usr/src/client

RUN apk update && apk upgrade
RUN apk add git

COPY . /usr/src/client/
RUN npm install
