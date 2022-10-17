#!/bin/bash

IMG="sunboy0213/http-server:latest"
BIN="server"

THIS_DIR=$(cd "$(dirname "$0")"; pwd)

GOOS=linux GOARCH=amd64 go build -o $THIS_DIR/$BIN $THIS_DIR/main.go

cd $THIS_DIR
docker build -t $IMG .

docker push $IMG

rm $BIN
