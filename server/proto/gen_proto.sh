#!/bin/sh

protoc --proto_path=. --go_out=.. --go-grpc_out=.. hello.proto
