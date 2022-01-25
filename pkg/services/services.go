package services

import (
	"github.com/eugeneuskov/grpc-chat/config"
	"github.com/eugeneuskov/grpc-chat/pkg/repositories"
)

type Services struct {
	External
	Auth
	authConfig *config.Auth
}

func NewServices(
	repositories *repositories.Repositories,
	authConfig *config.Auth,
) *Services {
	return &Services{
		External: newExternalService(repositories.External, authConfig.HashSalt),
		Auth:     newAuthService(repositories.Auth, authConfig),
	}
}
