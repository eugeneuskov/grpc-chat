package repositories

import (
	"fmt"
	"github.com/eugeneuskov/grpc-chat/pkg/structs"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type externalPostgres struct {
	db *gorm.DB
}

func (e *externalPostgres) CheckToken(token string) error {
	var externalAuth structs.ExternalAuth

	if err := e.db.Where("token = ? AND active = ?", token, true).First(&externalAuth); err.Error != nil {
		return fmt.Errorf("CheckToken error: %s", err.Error)
	}

	return nil
}

func (e *externalPostgres) CreateUser(user *structs.User) error {
	user.ID = uuid.NewV4()

	if err := e.db.Create(user); err.Error != nil {
		return fmt.Errorf("create user error: %s", err.Error)
	}

	return nil
}

func newExternalAuthPostgres(db *gorm.DB) *externalPostgres {
	return &externalPostgres{db}
}
