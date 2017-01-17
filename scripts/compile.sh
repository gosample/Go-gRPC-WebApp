#!/bin/bash

protoc -I=$GOPATH/src/stars-app/ \
-I$GOPATH/src \
$GOPATH/src/stars-app/messages/user/user.proto \
--go_out=plugins=grpc:$GOPATH/src/stars-app

protoc -I=$GOPATH/src/stars-app/ \
-I$GOPATH/src \
$GOPATH/src/stars-app/messages/ghResponse/ghList.proto \
--go_out=plugins=grpc:$GOPATH/src/stars-app

protoc -I=$GOPATH/src/stars-app/ \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
$GOPATH/src/stars-app/services/login/login.proto \
--go_out=plugins=grpc:$GOPATH/src/stars-app

protoc -I=$GOPATH/src/stars-app/ \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
$GOPATH/src/stars-app/services/list/list.proto \
--go_out=plugins=grpc:$GOPATH/src/stars-app


protoc -I=$GOPATH/src/stars-app/ \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--grpc-gateway_out=logtostderr=true:$GOPATH/src/stars-app/ \
$GOPATH/src/stars-app/services/login/login.proto

protoc -I=$GOPATH/src/stars-app/ \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--grpc-gateway_out=logtostderr=true:$GOPATH/src/stars-app/ \
$GOPATH/src/stars-app/services/list/list.proto
