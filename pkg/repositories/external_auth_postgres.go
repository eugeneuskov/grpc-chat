package repositories

import (
	"github.com/eugeneuskov/grpc-chat/pkg/structs"
	"gorm.io/gorm"
	"log"
)

type externalAuthPostgres struct {
	db *gorm.DB
}

func (e *externalAuthPostgres) CheckToken(token string) error {
	var externalAuth structs.ExternalAuth

	e.db.Where("token = ? AND active = ?", token, true).First(&externalAuth)

	log.Printf("%v\n", externalAuth)

	return nil
}

func (e *externalAuthPostgres) CreateUser() error {
	//TODO implement me
	panic("implement me")
}

func newExternalAuthPostgres(db *gorm.DB) *externalAuthPostgres {
	return &externalAuthPostgres{db: db}
}
