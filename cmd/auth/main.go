package main

import (
	"github.com/golang/protobuf/jsonpb"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"

	// "github.com/grpc-ecosystem/grpc-gateway/runtime"
	// "github.com/grpc-ecosystem/grpc-gateway/utilities"

	"os"

	cli "github.com/jawher/mow.cli"
	log "github.com/sirupsen/logrus"

	"github.com/wcrbrm/grpc-nginx-example/common/logging"
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
	m := jsonpb.Marshaler{}
	paginationJSON, _ := m.MarshalToString(&pb.Pagination{
		From:  &wrappers.StringValue{Value: "10"},
		Limit: &wrappers.Int64Value{Value: 50},
	})
	log.Printf("pagination: %v", paginationJSON)

	idJSON, _ := m.MarshalToString(&pb.ID{
		Value: "some-value",
	})
	log.Printf("id: %v", idJSON)

	strJSON, _ := m.MarshalToString(&wrappers.StringValue{
		Value: "some-string",
	})
	log.Printf("id: %v", strJSON)

	// lis, err := net.Listen("tcp", fmt.Sprintf(":%v", *appGrpcPort))
	// if err != nil {
	// 	log.WithError(err).Fatal("failed to listen")
	// }
	// s := grpc.NewServer()
	// pb.RegisterAuthServiceServer(s, auth.NewAuthService())
	// // Register reflection service on gRPC server.
	// reflection.Register(s)
	// log.WithField("addr", lis.Addr()).Info("Starting Server")
	// if err := s.Serve(lis); err != nil {
	// 	log.WithError(err).Fatalf("failed to serve")
	// }

}
