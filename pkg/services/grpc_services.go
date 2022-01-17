package services

import (
	"google.golang.org/grpc"
)

type GrpcServices struct {
	grpcServer *grpc.Server
}

func NewGrpcServices(grpcServer *grpc.Server) *GrpcServices {
	return &GrpcServices{
		grpcServer,
	}
}

func (s *GrpcServices) InitServices() {
	newGrpcBroadcastService(s.grpcServer)
	newGrpcExternalService(s.grpcServer)
}
