package repository

import (
	budget "budget-tracker"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user *budget.User) error
	GetUser(email, password string) (user budget.User, err error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: newAuthorizationSQL(db),
	}
}
