PROTOC_VERSION=3.9.1
PROTOC_ZIP=protoc-${PROTOC_VERSION}-linux-x86_64.zip
PROTOC_INCLUDES := -I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/ 

all:

.PHONY: proto auth up down deps redoc

proto:
	protoc -I . --go_out=plugins=grpc:. ./core/auth/pb/*.proto

swagger:
	protoc -I . $(PROTOC_INCLUDES) \
		--swagger_out=json_names_for_fields=true:. \
		./core/auth/pb/*.proto

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
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u google.golang.org/grpc
	sudo npm install -g redoc-cli

redoc: