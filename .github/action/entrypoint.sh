#!/bin/bash

set -exa

wget --retry-connrefused --no-check-certificate -T 60 https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz -P /usr/local/
tar -zxf /usr/local/go${GO_VERSION}.linux-amd64.tar.gz -C /usr/local
ln -s /usr/local/go/bin/go /usr/local/bin/go

go get -u github.com/jandelgado/gcov2lcov

golangci-lint run --timeout 10m

go test -race -coverprofile=coverage.out -covermode=atomic $(go list ./... | grep -v /cmd/)

~/go/bin/gcov2lcov -infile=coverage.out -outfile=coverage.lcov
