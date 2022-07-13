REVISION := $(shell git describe --always)
DATE := $(shell date +%Y-%m-%dT%H:%M:%S%z)
LDFLAGS	:= -ldflags="-X \"main.Revision=$(REVISION)\" -X \"main.BuildDate=${DATE}\" -extldflags \"-static\""

.PHONY: build-cross dist build mod clean run help docker

name		:= pstore
linux_name	:= $(name)-linux-amd64
darwin_name	:= $(name)-darwin-amd64
GO_VERSION      := 1.12

help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ { printf "\033[36m%-22s\033[0m %s\n", $$1, $$NF }' $(MAKEFILE_LIST)

dist: build-docker ## create .tar.gz linux & darwin to /bin
	cd bin && tar zcvf $(linux_name).tar.gz $(linux_name) && rm -f $(linux_name)
	cd bin && tar zcvf $(darwin_name).tar.gz $(darwin_name) && rm -f $(darwin_name)

build-cross: clean ## create to build for linux & darwin to bin/
	GOOS=linux GOARCH=amd64 go build -o bin/$(linux_name) $(LDFLAGS) *.go
	GOOS=darwin GOARCH=amd64 go build -o bin/$(darwin_name) $(LDFLAGS) *.go

build: ## go build
	go build -o bin/$(name) $(LDFLAGS) *.go

build-docker: ## go build on Docker
	@docker run --rm -v "$(PWD)":/go/src/github.com/sioncojp/$(name) -w /go/src/github.com/sioncojp/$(name) golang:$(GO_VERSION) bash build.sh

test: ## go test
	go test -v $$(go list ./... | grep -v /vendor/)

mod: ## go mod init
	go mod init

clean: ## remove bin/*
	rm -f bin/*

run: ## go run
	go run main.go

lint: ## go lint ignore vendor
	golint $(go list ./... | grep -v /vendor/)

docker: ## docker build
	docker build -t $(name) .
