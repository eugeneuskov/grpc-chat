package api

import (
	"context"
	"github.com/eugeneuskov/grpc-chat/pkg/services"
	"github.com/eugeneuskov/grpc-chat/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcAuthService struct {
	service services.Auth
}

func (g *grpcAuthService) Login(_ context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := g.service.Login(request.GetLogin(), request.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "wrong login or password")
	}

	return &pb.LoginResponse{Token: token}, nil
}

func (g *grpcAuthService) Info(_ context.Context, req *pb.InfoRequest) (*pb.InfoResponse, error) {
	user, err := g.service.CheckToken(req.GetToken()) // TODO На самом деле токен будет парситься из заголовка, а не в теле запроса
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return &pb.InfoResponse{
		Id:       user.ID.String(),
		Username: user.Username,
	}, nil
}

func newGrpcAuthService(s *grpc.Server, service services.Auth) {
	if s != nil {
		pb.RegisterAuthServer(s, &grpcAuthService{service})
	}
}
