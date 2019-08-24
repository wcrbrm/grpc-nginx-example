PROTOC_VERSION=3.9.1
PROTOC_ZIP=protoc-${PROTOC_VERSION}-linux-x86_64.zip

all:

.PHONY: proto auth run

proto:
	protoc -I . --go_out=plugins=grpc:. ./core/auth/pb/*.proto

auth:
	docker build -f docker/auth/Dockerfile . -t grpcexample-auth

up:
	docker-compose up -d --build

down:
	docker-compose down

build:
	go build -o example-auth github.com/wcrbrm/grpc-nginx-example/cmd/auth

deps:
	curl -OL https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP}
	sudo unzip -o ${PROTOC_ZIP} -d /usr/local bin/protoc
	sudo unzip -o ${PROTOC_ZIP} -d /usr/local include/*
	rm -f ${PROTOC_ZIP}
	go get -u github.com/golang/protobuf/protoc-gen-go