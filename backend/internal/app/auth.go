package app

import (
	budget "budget-tracker"
	"budget-tracker/pkg/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *App) signUp(c *gin.Context) {
	var input budget.User

	if err := c.BindJSON(&input); err != nil {
		app.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := app.s.CreateUser(&input)
	if err != nil {
		if errors.Is(err, repository.UniqueViolationError) {
			app.newErrorResponse(c, http.StatusConflict, "user is already exists")
			return
		}
		app.newErrorResponse(c, http.StatusInternalServerError, "something gone wrong while trying create user")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (app *App) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		app.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := app.s.GenerateToken(input.Email, input.Password)
	if err != nil {
		app.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
