package repositories

type ExternalAuth interface {
	CheckToken(token string) error
	CreateUser() error
}
