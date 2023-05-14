package repository

import (
	budget "budget-tracker"
	"gorm.io/gorm"
)

type authSql struct {
	db *gorm.DB
}

func newAuthorizationSQL(db *gorm.DB) Authorization {
	return &authSql{db: db}

}

func (db *authSql) CreateUser(user *budget.User) (err error) {
	tx := db.db.Begin()

	err = tx.Create(user).Error
	if err != nil {
		tx.Rollback()
		if handleError(err) != nil {
			return err
		}
		return UniqueViolationError
	}
	return tx.Commit().Error
}

func (db *authSql) GetUser(email, password string) (user budget.User, err error) {
	err = db.db.Model(&budget.User{Email: email, Password: password}).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
