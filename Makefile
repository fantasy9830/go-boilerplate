DIST := dist

FILENAME ?= goapp

GO ?= go
GOFMT ?= gofmt
GOLINT ?= golint
WIRE ?= wire

.PHONY: build clean lint fmt vet generate

build-dir:
	@mkdir -p $(DIST)

build: build-dir generate
	@$(GO) build -ldflags '-s -w' -o=$(DIST)/$(FILENAME)

clean:
	@rm -rf $(DIST)

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
