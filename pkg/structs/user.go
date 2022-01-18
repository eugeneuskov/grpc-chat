package structs

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	ExternalId string    `gorm:"size:255;index:idx_external_id,unique" json:"external_id"`
	Login      string    `gorm:"size:255;index:idx_login,unique" json:"login"`
	Password   string    `gorm:"size:255" json:"password"`
	Username   string    `gorm:"size:255" json:"username"`
}

func (u *User) IsEmpty() bool {
	return u.ExternalId == "" ||
		u.Login == "" ||
		u.Password == "" ||
		u.Username == ""
}
