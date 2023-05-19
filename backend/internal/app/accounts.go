package app

import (
	budget "budget-tracker"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *App) createAccount(c *gin.Context) {
	var input budget.Account

	userId, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = c.BindJSON(&input); err != nil {
		app.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.UserID = userId

	err = app.s.AccountsService.CreateAccount(input)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func (app *App) getAllAccounts(c *gin.Context) {

	userId, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	orderBy := c.Query("orderBy")
	if orderBy == "" {
		orderBy = "date"
	}

	account := c.Param("account")
	if account != "deposit" && account != "credit" {
		account = ""
	}

	sortedBy := c.Query("reverse")
	if sortedBy == "" || sortedBy == "false" {
		sortedBy = " desc"
	} else if sortedBy == "true" {
		sortedBy = " asc"
	} else {
		app.newErrorResponse(c, http.StatusBadRequest, "invalid query param reverse")
		return
	}

	list, err := app.s.AccountsService.GetAll(userId, account, orderBy, sortedBy)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": list,
	})

}

func (app *App) deleteAccount(c *gin.Context) {
	var input deleteInput

	userId, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = c.BindJSON(&input); err != nil {
		app.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = app.s.AccountsService.DeleteAccount(input.Id, userId)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})

}

func (app *App) editAccount(c *gin.Context) {
	var input budget.Account

	userId, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = c.BindJSON(&input); err != nil {
		app.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.UserID = userId

	err = app.s.AccountsService.EditAccount(input)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, input)

}

func (app *App) getTotalDeposits(c *gin.Context) {
	userId, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	payed, goalSum := app.s.AccountsService.GetTotalsDeposits(userId)

	c.JSON(http.StatusOK, map[string]interface{}{
		"payed":   payed,
		"goalSum": goalSum,
	})

}
