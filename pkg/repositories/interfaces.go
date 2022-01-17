package repositories

type ExternalAuth interface {
	CheckToken() error
	CreateUser() error
}
