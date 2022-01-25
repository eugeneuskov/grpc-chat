package services

import (
	"github.com/eugeneuskov/grpc-chat/pkg/repositories"
	"github.com/eugeneuskov/grpc-chat/pkg/structs"
)

type externalService struct {
	repository repositories.External
	hashSalt   string
}

func (e *externalService) CheckToken(token string) error {
	return e.repository.CheckToken(token)
}

func (e *externalService) CreateUser(user *structs.User) error {
	user.Password = generatePasswordHash(user.Password, e.hashSalt)
	return e.repository.CreateUser(user)
}

func newExternalService(
	repository repositories.External,
	hashSalt string,
) *externalService {
	return &externalService{
		repository,
		hashSalt,
	}
}
