package auth

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/wcrbrm/grpc-nginx-example/common/logging"
	pb "github.com/wcrbrm/grpc-nginx-example/core/auth/pb"
)

// authServiceServerImpl is the actual of auth.AuthServiceServer
type authServiceServerImpl struct {
	fields log.Fields
}

// NewAuthService returns an implement
func NewAuthService() pb.AuthServiceServer {
	return &authServiceServerImpl{
		fields: log.Fields{"service": "auth"},
	}
}

func (s *authServiceServerImpl) AccountAuth(ctx context.Context, in *pb.AuthRequest) (*pb.AuthResponse, error) {
	ID := pb.ID{Value: "RANDOM_ID"}
	Result := fmt.Sprintf("Hello %v, your ID is %v", in.Username, ID)

	log.WithFields(logging.WithFn(log.Fields{"username": in.Username, "code": in.Code})).
		Info(Result)
	return &pb.AuthResponse{Result: Result}, nil
}
