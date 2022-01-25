package repositories

import "github.com/eugeneuskov/grpc-chat/pkg/structs"

type External interface {
	CheckToken(token string) error
	CreateUser(user *structs.User) error
}

type Auth interface {
	Login(login, password string) (*structs.User, error)
}
