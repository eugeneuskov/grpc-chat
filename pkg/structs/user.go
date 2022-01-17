package structs

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	ExternalId string    `gorm:"size:255;index:idx_external_id,unique"`
	Login      string    `gorm:"size:255;index:idx_login,unique"`
	Password   string    `gorm:"size:255"`
	Username   string    `gorm:"size:255"`
}
