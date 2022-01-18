package workers

import "github.com/eugeneuskov/grpc-chat/pkg/structs"

type Worker interface {
	Start()
}

type CreateUser interface {
	Worker
	validate(message []byte) (*structs.User, error)
}
