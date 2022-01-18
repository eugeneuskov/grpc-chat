package api

import (
	"github.com/eugeneuskov/grpc-chat/pkg/services"
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

func (s *GrpcServices) InitServices(serviceList *services.Services) {
	newGrpcBroadcastService(s.grpcServer)
	newGrpcExternalService(s.grpcServer, serviceList.ExternalAuth)
}
