#!/usr/bin/env bash

proto_dir=$(ls -d api/proto/*)

for dir in ${proto_dir[@]}; do
  echo "==> Generate ${dir} Protocol Buffers..."

  protoc --proto_path=${dir} --go_out=${dir} --go-grpc_out=${dir} ${dir}/*.proto
done;
