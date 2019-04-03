#!/bin/bash

mode=${1:-local}
echo "Run Mode=$mode"

pkill builderTools
export GOPATH=/Users/saisatch/myWorkspace/builder-factory
rm -Rf build
mkdir build
cp config.dev.yaml build/
cp config.prod.yaml build/

if [ $mode = "docker" ]
then
    GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o build/builderTools main.go 
    docker build -t auto-devops-app-service-rpc .
    docker run --rm -it -p 50051:50051 -p 8080:8080 --name grpc_service auto-devops-app-service-rpc
else
    export RUN_ENV=dev
    go build -o build/builderTools main.go 
    ./build/builderTools & ./grpcwebproxy --run_tls_server=false --allow_all_origins --backend_addr=localhost:50051
fi