package services

type ExternalAuth interface {
	CheckToken() error
	CreateUser() error
}
