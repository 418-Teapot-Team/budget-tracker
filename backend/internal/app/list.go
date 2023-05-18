package app

import (
	budget "budget-tracker"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *App) createList(c *gin.Context) {
	var input budget.List

	userId, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = c.BindJSON(&input); err != nil {
		app.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err = app.s.ListsService.CreateList(&input, userId)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

type deleteInput struct {
	Id int `json:"id"`
}

func (app *App) deleteList(c *gin.Context) {
	var input deleteInput

	userId, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = c.BindJSON(&input); err != nil {
		app.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err = app.s.ListsService.DeleteList(input.Id, userId)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})

}

func (app *App) getBudgetList(c *gin.Context) {

	orderBy := c.Query("orderBy")
	if orderBy == "" {
		orderBy = "created_at"
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

	budgetType := c.Param("type")
	if budgetType == "" {
		app.newErrorResponse(c, http.StatusBadRequest, "please indicate the type of budget")
		return
	}

	if budgetType != "income" && budgetType != "expences" {
		budgetType = ""
	}
	userId, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list, _ := app.s.GetList(userId, budgetType, orderBy, sortedBy)

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": list,
	})
}

func (app *App) editBudgetList(c *gin.Context) {
	var input budget.List

	userid, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = c.BindJSON(&input); err != nil {
		app.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	input.UserId = userid

	err = app.s.EditList(input)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, input)

}

func (app *App) getTopExpenses(c *gin.Context) {

	userId, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list, err := app.s.GetTopExpenses(userId)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": list,
	})
}
