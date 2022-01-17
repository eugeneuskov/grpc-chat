package services

import (
	"context"
	"github.com/eugeneuskov/grpc-chat/pkg/server"
	"github.com/eugeneuskov/grpc-chat/proto/pb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const authHeaderName = "X-Access-Token"

type grpcExternalService struct {
	service ExternalAuth
}

func (e *grpcExternalService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*emptypb.Empty, error) {
	token, err := server.GetTokenByName(ctx, authHeaderName)
	if err != nil {
		return nil, err
	}

	if err = e.service.CheckToken(token); err != nil {
		return nil, status.Error(codes.Unauthenticated, "wrong token")
	}

	return &empty.Empty{}, nil
}

func newGrpcExternalService(s *grpc.Server, service ExternalAuth) {
	if s != nil {
		pb.RegisterExternalServer(s, &grpcExternalService{service: service})
	}
}
