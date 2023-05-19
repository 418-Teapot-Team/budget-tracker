package service

import (
	budget "budget-tracker"
	"budget-tracker/pkg/repository"
	"errors"
	"fmt"
	"time"
)

const layout = "2006-01-02"

type AccountsService struct {
	repo repository.Accounts
}

func (as *AccountsService) CreateAccount(input budget.Account) (err error) {
	_struct, currMonth := getUpdatedStruct(input)

	date, _ := time.Parse(layout, input.Date)

	if date.After(time.Now()) {
		return errors.New(fmt.Sprintf("Error occurred, %s time cannot be bigger than current time", input.Type))

	}

	if currMonth > _struct.MonthAmount {
		return errors.New(fmt.Sprintf("wrong date range, impossible to create this %s", input.Type))
	}

	return as.repo.CreateAccount(&_struct)
}

type listOutput struct {
	Id           int        `json:"id"`
	Type         string     `json:"type"`
	Name         string     `json:"name"`
	MonthAmount  int        `json:"monthAmount"`
	CurrentMonth int        `json:"currentMonth"`
	Percent      float64    `json:"percent"`
	Sum          float64    `json:"sum"`
	MonthPayment float64    `json:"monthPayment"`
	Payed        float64    `json:"payed"`
	Date         string     `json:"startAccountDate"`
	FinishDate   string     `json:"finishAccountDate"`
	GoalSum      float64    `json:"goalSum"`
	CreatedAt    *time.Time `json:"createdAt" binding:"required"`
}

func (as *AccountsService) GetAll(userId int, account, orderBy, sortedBy string) ([]listOutput, error) {
	var outputList []listOutput

	list, err := as.repo.GetAllAccounts(userId, account, orderBy, sortedBy)
	if err != nil {
		return outputList, err
	}

	for _, finance := range list {

		_str, curMonth := getUpdatedStruct(finance)

		outputList = append(outputList, listOutput{
			Id:           finance.ID,
			Type:         finance.Type,
			Name:         finance.Name,
			MonthAmount:  finance.MonthAmount,
			CurrentMonth: curMonth,
			Percent:      finance.Percent,
			Sum:          finance.Sum,
			MonthPayment: _str.MonthlyPayment,
			Payed:        float64(curMonth) * _str.MonthlyPayment,
			Date:         finance.Date,
			FinishDate:   _str.FinishDate,
			GoalSum:      _str.GoalSum,
			CreatedAt:    &_str.CreatedAt,
		})
	}

	return outputList, nil
}

func getDaysFromDate(date string) int {
	targetDate, err := time.Parse(layout, date[:len(layout)])
	if err != nil {
		fmt.Println("Error parsing target date:", err)
		return 0
	}
	currentDate := time.Now().UTC().Truncate(24 * time.Hour)
	duration := currentDate.Sub(targetDate)
	days := uint(duration.Hours() / 24)
	return int(days)

}

func (as *AccountsService) DeleteAccount(listId, userId int) (err error) {
	return as.repo.DeleteAccount(listId, userId)
}

func (as *AccountsService) EditAccount(finance budget.Account) (err error) {
	_str, currMonth := getUpdatedStruct(finance)
	if currMonth > _str.MonthAmount {
		return errors.New(fmt.Sprintf("wrong date range, impossible to update this %s", finance.Type))
	}

	return as.repo.EditAccount(_str)
}

func getUpdatedStruct(finance budget.Account) (budget.Account, int) {
	var mounthPayment float64

	date, _ := time.Parse(layout, finance.Date)
	finishDateStr := date.Add(time.Hour * 24 * 30 * time.Duration(finance.MonthAmount)).String()

	currentDate, _ := time.Parse(layout, time.Now().Format(layout))

	duration := currentDate.Sub(date)

	currentMonthInt := int(duration.Hours() / (24 * 30)) //

	if finance.Type == "deposit" {
		mounthPayment = finance.Sum * (finance.Percent / 100) / 12
	} else if finance.Type == "credit" {
		daysOfCredit := getDaysFromDate(finance.Date) // 2023-01-11 // 127

		mounthPayment = (finance.Sum*(finance.Percent/100)*float64(daysOfCredit)/365 + finance.Sum) / float64(finance.MonthAmount)
	}

	var (
		goalSum         float64
		alreadyReceived float64
	)
	if finance.Type == "deposit" {
		goalSum = mounthPayment*float64(finance.MonthAmount) + finance.Sum
	} else {
		goalSum = mounthPayment * float64(finance.MonthAmount)
	}

	if finance.Type == "deposit" {
		alreadyReceived = float64(currentMonthInt) * (mounthPayment + finance.Sum/12)
	} else {
		alreadyReceived = float64(currentMonthInt) * mounthPayment
	}

	update := budget.Account{
		ID:              finance.ID,
		Type:            finance.Type,
		UserID:          finance.UserID,
		Name:            finance.Name,
		MonthAmount:     finance.MonthAmount,
		Percent:         finance.Percent,
		Date:            finance.Date,
		AlreadyReceived: alreadyReceived,
		FinishDate:      finishDateStr[:len(layout)],
		MonthlyPayment:  mounthPayment,
		Sum:             finance.Sum,
		GoalSum:         goalSum,
		CreatedAt:       finance.CreatedAt,
	}
	return update, currentMonthInt
}

func (as *AccountsService) GetTotalsDeposits(userId int) (float64, float64) {
	return as.repo.GetTotalsDeposits(userId)
}
