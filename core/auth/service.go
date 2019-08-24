package auth

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/wcrbrm/grpc-nginx-example/common/logging"
	"github.com/wcrbrm/grpc-nginx-example/core/auth/pb"
)

// authServiceServerImpl is the actual of pb.AuthServiceServer
type authServiceServerImpl struct {
	pb.AuthServiceServer
	fields log.Fields
}

// NewAuthService returns an implement
func NewAuthService() pb.AuthServiceServer {
	return &authServiceServerImpl{
		fields: log.Fields{"service": "auth"},
	}
}

func (s *authServiceServerImpl) AccountAuth(ctx context.Context, in *pb.AuthRequest) (*pb.AuthResponse, error) {
	log.WithFields(logging.WithFn(log.Fields{"username": in.Username, "code": in.Code})).Info("auth")
	return &pb.AuthResponse{Message: "Hello " + in.Username}, nil
}
