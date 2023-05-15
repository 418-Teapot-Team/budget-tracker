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
