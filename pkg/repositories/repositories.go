package repositories

import "gorm.io/gorm"

type Repositories struct {
	ExternalAuth
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		ExternalAuth: newExternalAuthPostgres(db),
	}
}
