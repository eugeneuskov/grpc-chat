package api

import (
	"context"
	"github.com/eugeneuskov/grpc-chat/pkg/server"
	"github.com/eugeneuskov/grpc-chat/pkg/services"
	"github.com/eugeneuskov/grpc-chat/pkg/structs"
	"github.com/eugeneuskov/grpc-chat/proto/pb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const authHeaderName = "X-Access-Token"

type grpcExternalService struct {
	service services.ExternalAuth
}

func (e *grpcExternalService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*emptypb.Empty, error) {
	token, err := server.GetTokenByName(ctx, authHeaderName)
	if err != nil {
		return nil, err
	}

	if err = e.service.CheckToken(token); err != nil {
		return nil, status.Error(codes.Unauthenticated, "wrong token")
	}

	if err := e.service.CreateUser(createUserStruct(req)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func createUserStruct(req *pb.CreateUserRequest) *structs.User {
	return &structs.User{
		ExternalId: req.GetExternalId(),
		Login:      req.GetLogin(),
		Password:   req.GetPassword(),
		Username:   req.GetUsername(),
	}
}

func newGrpcExternalService(s *grpc.Server, service services.ExternalAuth) {
	if s != nil {
		pb.RegisterExternalServer(s, &grpcExternalService{service: service})
	}
}
