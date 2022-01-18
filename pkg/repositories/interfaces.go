package repositories

import "github.com/eugeneuskov/grpc-chat/pkg/structs"

type ExternalAuth interface {
	CheckToken(token string) error
	CreateUser(user *structs.User) error
}
