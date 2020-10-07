RELEASE := release
PROTO_DIR := api/protobuf

FILENAME ?= goapp
GOFILES := $(shell find . -name "*.go" -type f | grep -v "/vendor/")

GO ?= go
GOFMT ?= gofmt
GOLINT ?= golint
WIRE ?= wire
PROTOC ?= protoc

.PHONY: all build clean lint fmt vet generate proto wire

all: proto wire fmt lint vet build

build-dir:
	@mkdir -p $(RELEASE)

build: clean build-dir generate
	@$(GO) build -ldflags '-s -w' -o=$(RELEASE)/$(FILENAME)

clean:
	@rm -rf $(RELEASE)

generate:
	@$(GO) generate

lint:
	@$(GOLINT) $(GOFILES)

fmt:
	@$(GOFMT) -s -l -w $(GOFILES)

vet:
	@$(GO) vet ./...

wire:
	@$(WIRE) ./...

proto:
	@$(PROTOC) --go_out=$(PROTO_DIR) --go-grpc_out=$(PROTO_DIR) $(PROTO_DIR)/*.proto
