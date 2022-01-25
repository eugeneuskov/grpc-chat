package repositories

import "gorm.io/gorm"

type Repositories struct {
	External
	Auth
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		External: newExternalAuthPostgres(db),
		Auth:     newAuthPostgres(db),
	}
}
