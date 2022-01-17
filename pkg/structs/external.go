package structs

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ExternalAuth struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;primary_key"`
	Name   string    `gorm:"size:255"`
	Token  string    `gorm:"size:255;index:idx_token,unique"`
	Active bool
}
