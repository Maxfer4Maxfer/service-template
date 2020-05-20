#!/usr/bin/env sh

# go get -u google.golang.org/grpc
# go get -u github.com/golang/protobuf/protoc-gen-go
#
# Update protoc Go bindings via
#  go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
#
# See also
#  https://github.com/grpc/grpc-go/tree/master/examples

protoc calculator.proto --go_out=plugins=grpc:../../internal/transport/grpc/pb
