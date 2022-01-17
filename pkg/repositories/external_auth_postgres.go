package repositories

import "gorm.io/gorm"

type externalAuthPostgres struct {
	db *gorm.DB
}

func (e *externalAuthPostgres) CheckToken() error {
	//TODO implement me
	panic("implement me")
}

func (e *externalAuthPostgres) CreateUser() error {
	//TODO implement me
	panic("implement me")
}

func newExternalAuthPostgres(db *gorm.DB) *externalAuthPostgres {
	return &externalAuthPostgres{db: db}
}
