#!/bin/sh

protoc --proto_path=../proto --go_out=. --go-grpc_out=. hello.proto && \
go build -o ./jjgame main.go