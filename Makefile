PROTOC_VERSION=3.9.1
PROTOC_ZIP=protoc-${PROTOC_VERSION}-linux-x86_64.zip
PROTOC_PATH= \
    --proto_path=. \
	--proto_path=$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--proto_path=$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/ \
	--proto_path=$$GOPATH/src/github.com/amsokol/

PROTOC_PLUGINS=grpc

all:

.PHONY: proto auth up down deps redoc

proto:
	protoc $(PROTOC_PATH) \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=json_names_for_fields=true:. \
		--go_out=plugins=$(PROTOC_PLUGINS):. \
		./core/auth/pb/*.proto
	protoc $(PROTOC_PATH) \
		--gotag_out=xxx="bson+\"-\"",output_path=./core/auth/pb:.\
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
	go get -u github.com/amsokol/protoc-gen-gotag
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u google.golang.org/grpc
	go get -u github.com/amsokol/mongo-go-driver-protobuf
	go get -u github.com/amsokol/protoc-gen-gotag
	sudo npm install -g redoc-cli

redoc: