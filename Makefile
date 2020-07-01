RELEASE := release
PROTO_DIR := api/proto

FILENAME ?= goapp

GO ?= go
GOFMT ?= gofmt
GOLINT ?= golint
WIRE ?= wire
PROTOC ?= protoc

.PHONY: all build clean lint fmt vet generate proto wire

all: proto wire lint fmt vet build

build-dir:
	@mkdir -p $(RELEASE)

build: build-dir generate
	@$(GO) build -ldflags '-s -w' -o=$(RELEASE)/$(FILENAME)

clean:
	@rm -rf $(RELEASE)

generate:
	@$(GO) generate

lint:
	@$(GOLINT) ./...

fmt:
	@$(GOFMT) -s -l -w .

vet:
	@$(GO) vet ./...

wire:
	@$(WIRE) ./...

proto:
	@$(PROTOC) --go_out=$(PROTO_DIR) $(PROTO_DIR)/*.proto
