FROM golang:1.10.0-alpine3.7 AS build

COPY . /data/api/

RUN apk -U --no-cache add git protobuf

RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go

RUN mkdir -p /go/src/api
RUN protoc -I /data/ --go_out=plugins=grpc:/go/src/ /data/api/core/*.proto
RUN protoc -I /data/ --go_out=plugins=grpc:/go/src/ /data/api/people/*.proto
RUN protoc -I /data/ --go_out=plugins=grpc:/go/src/ /data/api/greeter/*.proto