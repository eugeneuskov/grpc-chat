package repositories

import (
	"fmt"
	"github.com/eugeneuskov/grpc-chat/pkg/structs"
	"gorm.io/gorm"
)

type authPostgres struct {
	db *gorm.DB
}

func (a *authPostgres) Login(login, password string) (*structs.User, error) {
	var user structs.User

	if err := a.db.Where("login = ? AND password = ?", login, password).First(&user); err.Error != nil {
		return nil, fmt.Errorf("user was not found")
	}

	return &user, nil
}

func newAuthPostgres(db *gorm.DB) *authPostgres {
	return &authPostgres{db}
}
