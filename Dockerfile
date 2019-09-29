FROM golang:1.13.0-buster

USER root

WORKDIR /go/src

RUN apt-get update && apt-get upgrade -y
RUN apt-get install -y vim less

RUN go get -u \
  google.golang.org/grpc \
  github.com/golang/protobuf/protoc-gen-go \
  github.com/kujtimiihoxha/kit

CMD ["/bin/sh"]
