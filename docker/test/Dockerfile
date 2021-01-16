FROM ubuntu:20.04

COPY ./docker/etc/sources.list /etc/apt

ARG DEBIAN_FRONTEND=noninteractive

RUN apt update && apt upgrade -y
RUN apt-get -y install libgl1-mesa-dev xorg-dev gcc
RUN apt-get -y install wget curl

ARG GO_VERSION=1.15.6

ADD https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz /usr/local/
RUN tar -zxf /usr/local/go${GO_VERSION}.linux-amd64.tar.gz -C /usr/local
RUN ln -s /usr/local/go/bin/go /usr/local/bin/go

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.31.0

RUN mkdir -p /go/src/app/
WORKDIR /go/src/app

COPY ./ /go/src/app

RUN golangci-lint run --timeout 10m
RUN go test -race $(go list ./... | grep -v /cmd/)
