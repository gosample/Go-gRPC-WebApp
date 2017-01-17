FROM golang

RUN go get google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/proto
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

COPY ../ /go/src/stars-app

EXPOSE 8587
