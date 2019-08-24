package main

import (
	"fmt"
	"net"
	"os"

	cli "github.com/jawher/mow.cli"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/wcrbrm/grpc-nginx-example/common/logging"
	"github.com/wcrbrm/grpc-nginx-example/core/auth"
	pb "github.com/wcrbrm/grpc-nginx-example/core/auth/pb"
)

var app = cli.App("auth", "example gRPC service for authentication")

func main() {
	log.SetLevel(logging.Level(*appLogLevel))
	log.SetFormatter(logging.DefaultFormatter())

	app.Action = startgRPCServer
	if err := app.Run(os.Args); err != nil {
		log.Fatalln("[ERR]", err)
	}
}

func startgRPCServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", *appGrpcPort))
	if err != nil {
		log.WithError(err).Fatal("failed to listen")
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, auth.NewAuthService())

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.WithError(err).Fatalf("failed to serve")
	}
	log.Info("Server started")
}
