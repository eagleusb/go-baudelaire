SHELL = /bin/bash

GOOS   ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOPATH ?= $(shell go env GOPATH)

export GOPATH

.PHONY = environment build dependency tests
.RECIPEPREFIX = >

VERSION := "v0.1.0"

environment:
> @echo -e "system environment:\n"
> @env | sort -u;
> @echo -e "\ngo environment:\n"
> @go env

dependency:
> go get -v -u ./...
> go mod tidy

lint:
> go vet -x ./...
> gofmt -d -s .

build:
> @echo -e "\nUsing $(GOPATH) as GOPATH";
> go build -i -v -o bin/baudelaire \
>  -ldflags="-s -w -X main.version=$(VERSION)" .;

tests:
> @echo -e "\nTesting using $(GOPATH) as GOPATH";
> go test -v ./...
