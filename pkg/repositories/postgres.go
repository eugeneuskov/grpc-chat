package repositories

import (
	"fmt"
	"github.com/eugeneuskov/grpc-chat/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(dbConfig *config.Database) (*gorm.DB, error) {
	return gorm.Open(
		postgres.Open(fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			dbConfig.Host,
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Name,
			dbConfig.Port,
			dbConfig.SslMode,
		)),
		&gorm.Config{},
	)
}
