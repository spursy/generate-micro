FROM docker.longbridge-inc.com/long-bridge-core-system/golang-builder:1.16.4

RUN apk add --update build-base

ENV GOPROXY=https://goproxy.dev.longbridge-inc.com

RUN mkdir /tmp/app && \
    apk add curl

COPY . /tmp/app/

RUN cd /tmp/app && \
go mod vendor
