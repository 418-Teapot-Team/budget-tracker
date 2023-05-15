package repository

import (
	budget "budget-tracker"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user *budget.User) error
	GetUserAuth(email, password string) (user budget.User, err error)
	GetUserById(userId int) (user budget.User, err error)
}

type Lists interface {
	CreateList(input *budget.List) (err error)
	DeleteList(listId, userId int) (err error)
}

type Repository struct {
	Authorization
	Lists
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: newAuthorizationSQL(db),
		Lists:         newListsSQL(db),
	}
}
