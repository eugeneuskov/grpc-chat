package services

import (
	"github.com/eugeneuskov/grpc-chat/config"
	"github.com/eugeneuskov/grpc-chat/pkg/repositories"
)

type Services struct {
	ExternalAuth
	authConfig *config.Auth
}

func NewServices(
	repositories *repositories.Repositories,
	authConfig *config.Auth,
) *Services {
	return &Services{
		ExternalAuth: newExternalService(repositories.ExternalAuth, authConfig.HashSalt),
	}
}
