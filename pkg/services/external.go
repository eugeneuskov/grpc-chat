package services

import "github.com/eugeneuskov/grpc-chat/pkg/repositories"

type externalService struct {
	repository repositories.ExternalAuth
}

func (e *externalService) CheckToken() error {
	//TODO implement me
	panic("implement me")
}

func (e *externalService) CreateUser() error {
	//TODO implement me
	panic("implement me")
}

func newExternalService(repository repositories.ExternalAuth) *externalService {
	return &externalService{repository: repository}
}
