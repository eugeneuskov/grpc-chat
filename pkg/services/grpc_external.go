package services

import (
	"context"
	"github.com/eugeneuskov/grpc-chat/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type grpcExternalService struct{}

func (e *grpcExternalService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func newGrpcExternalService(s *grpc.Server) {
	if s != nil {
		pb.RegisterExternalServer(s, &grpcExternalService{})
	}
}
