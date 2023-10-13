GO ?= go
WIRE ?= wire

PACKAGES ?= $(shell $(GO) list ./... | grep -v /vendor/)

.PHONY: build check fmt generate proto docs

# 編譯所有的服務
build:
	@sh -c "'$(CURDIR)/scripts/build.sh'"

# 相關檢查
check: fmt
	@sh -c "'$(CURDIR)/scripts/check.sh'"

# 格式化文件
fmt:
	@sh -c "'$(CURDIR)/scripts/fmt.sh'"

# test
test:
	@$(GO) test -race -cover $(PACKAGES)

# 執行 go generate
generate:
	@$(GO) generate ./...

# 執行 wire
wire:
	@$(WIRE) ./...

# proto 檔轉成 go 程式
proto:
	@sh -c "'$(CURDIR)/scripts/proto.sh'"
