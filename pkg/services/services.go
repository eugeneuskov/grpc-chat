package services

import "github.com/eugeneuskov/grpc-chat/pkg/repositories"

type Services struct {
	ExternalAuth
}

func NewServices(repositories *repositories.Repositories) *Services {
	return &Services{
		ExternalAuth: newExternalService(repositories.ExternalAuth),
	}
}
