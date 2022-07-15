FROM golang:1.18-alpine

RUN apk update && apk --no-cache --update add build-base
RUN go get -u google.golang.org/grpc

RUN go get -u google.golang.org/protobuf/cmd/protoc-gen-go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go

RUN go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc