FROM golang:1.12
WORKDIR /apps
ADD go.mod go.sum ./
RUN go mod download
ADD common/   /apps/common/ 
ADD cmd/      /apps/cmd/
ADD core/     /apps/core/
RUN go build -o /apps/bin/auth github.com/wcrbrm/grpc-nginx-example/cmd/auth

FROM alpine:3.9.4
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=0 /apps/bin/auth /usr/bin/auth
ENTRYPOINT ["/usr/bin/auth"]