#! /bin/bash
# set Env path
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
# Install protoc
go get -u google.golang.org/grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
echo -e "\033[42;33m>>>\033[0m [BEGIN]"
# Trailer
echo ">>> Generating Trailer Proto..."
protoc -I ./ --go_out ./ --go_opt paths=source_relative \
    --go-grpc_out=./ --go-grpc_opt paths=source_relative \
    ./trailer.proto
echo ">>> Generate Trailer Proto OK."

echo -e "\033[42;33m>>>\033[0m [FINISHED]"
