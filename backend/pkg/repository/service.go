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
	GetTopCategories(userId int, enum string) (lists []budget.ListsGetter, err error)
	EditList(input budget.List) (err error)
	GetList(userId int, budgetType, orderBy, sortedBy string, takeAmount int, skipAmount int, categoryId int) (lists []budget.ListsGetter, err error)
	GetCurrentMonthSavings(userId int) (result float64, err error)
	GetSavingsStats(userId int) (data []budget.FinancialData, err error)
	GetTotalAmount(userId int, lType string, months int) (result float64, err error)
	GetStats(userId int, lType string, months int) (data []budget.FinancialData, err error)
}

type Categories interface {
	GetAllCategories() (categories []budget.Categories, err error)
}

type Accounts interface {
	CreateAccount(input *budget.Account) (err error)
	GetAllAccounts(userId int, account, orderBy, sortedBy string) (list []budget.Account, err error)
	DeleteAccount(listId, userId int) (err error)
	EditAccount(input budget.Account) (err error)
	GetTotalsDeposits(userId int) (float64, float64)
}

type Repository struct {
	Authorization
	Lists
	Categories
	Accounts
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: newAuthorizationSQL(db),
		Lists:         newListsSQL(db),
		Categories:    newCategorySQl(db),
		Accounts:      newAccountsSQL(db),
	}
}
