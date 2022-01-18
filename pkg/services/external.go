package services

import (
	"crypto/sha1"
	"fmt"
	"github.com/eugeneuskov/grpc-chat/pkg/repositories"
	"github.com/eugeneuskov/grpc-chat/pkg/structs"
)

type externalService struct {
	repository repositories.ExternalAuth
	hashSalt   string
}

func (e *externalService) CheckToken(token string) error {
	return e.repository.CheckToken(token)
}

func (e *externalService) CreateUser(user *structs.User) error {
	user.Password = e.generatePasswordHash(user.Password)
	return e.repository.CreateUser(user)
}

func (e *externalService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(e.hashSalt)))
}

func newExternalService(
	repository repositories.ExternalAuth,
	hashSalt string,
) *externalService {
	return &externalService{
		repository,
		hashSalt,
	}
}
