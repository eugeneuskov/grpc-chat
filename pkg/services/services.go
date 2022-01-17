package services

import (
	"google.golang.org/grpc"
)

type Service struct {
	grpcServer *grpc.Server
}

func NewService(grpcServer *grpc.Server) *Service {
	return &Service{
		grpcServer,
	}
}

func (s *Service) InitServices() {
	newBroadcastService(s.grpcServer)
	newExternalService(s.grpcServer)
}
