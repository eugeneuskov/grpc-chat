package server

import (
	"fmt"
	"github.com/eugeneuskov/grpc-chat/config"
	"github.com/eugeneuskov/grpc-chat/pkg/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

type Server struct {
	Grpc      *grpc.Server
	appConfig *config.App
	tlsConfig *config.Tls
}

func NewServer(tlsConfig *config.Tls, appConfig *config.App) *Server {
	return &Server{
		tlsConfig: tlsConfig,
		appConfig: appConfig,
	}
}

func (s *Server) Run() {
	s.Grpc = grpc.NewServer(s.options()...)

	println(fmt.Sprintf("Server running at %s port", s.appConfig.Port))
	lis, err := net.Listen("tcp", s.appConfig.Port)
	if err != nil {
		log.Fatalf("Failed listen: %s\n", err.Error())
	}

	services.NewService(s.Grpc).InitServices()

	if err = s.Grpc.Serve(lis); err != nil {
		log.Fatalf("Error occured while running gRPC HTTP2 server: %s\n", err.Error())
	}
}

func (s *Server) options() []grpc.ServerOption {
	opts := make([]grpc.ServerOption, 0, 4)

	opts = append(
		opts,
		grpc.MaxSendMsgSize(5*1024*1024*1024*1024),
		grpc.MaxRecvMsgSize(5*1024*1024*1024*1024),
	)

	if s.tlsConfig.Mode {
		creds, err := credentials.NewServerTLSFromFile(s.tlsConfig.CertFile, s.tlsConfig.KeyFile)
		if err != nil {
			log.Fatalf("Failed loading TLS: %s\n", err.Error())
			return nil
		}

		opts = append(opts, grpc.Creds(creds))
	}

	return opts
}
