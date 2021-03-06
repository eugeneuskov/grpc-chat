package repositories

import (
	"github.com/eugeneuskov/grpc-chat/pkg/structs"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(structs.User{}, structs.ExternalAuth{}); err != nil {
		return err
	}

	return nil
}
