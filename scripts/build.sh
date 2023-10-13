#!/usr/bin/env bash

export CGO_ENABLED=0

if [[ -z "${VERSION}" ]]; then
  VERSION=develop
fi

BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
EXTLDFLAGS='-extldflags "-static -fPIC"'
LDFLAGS="-s -w -X go-boilerplate/pkg/version.Version=${VERSION} -X go-boilerplate/pkg/version.BuildDate=${BUILD_DATE}"

# Delete the old dir
echo "==> Removing old directory..."
rm -rf release/*
mkdir -p release/

# Ensure all remote modules are downloaded and cached
go mod download

echo "==> Building..."
go build -ldflags "${LDFLAGS} ${EXTLDFLAGS}" -o release ./cmd/*

# Done!
echo "==> Results:"
ls -hl release/
