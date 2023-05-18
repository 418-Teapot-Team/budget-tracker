package app

import (
	budget "budget-tracker"
	"budget-tracker/pkg/repository"
	"errors"
	"fmt"
	"github.com/BoryslavGlov/logrusx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *App) signUp(c *gin.Context) {
	var input budget.User

	app.logg.Info("User trying to sign-in", logrusx.LogField{Key: "request", Value: fmt.Sprintf("%+v", c.Request)})

	if err := c.BindJSON(&input); err != nil {
		app.newErrorResponse(c, http.StatusBadRequest, err.Error())
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

func (app *App) whoAmI(c *gin.Context) {
	userId, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	user, err := app.s.AuthService.GetUserById(userId)
	if err != nil || user.Id == 0 {
		app.newErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"userId":   user.Id,
		"fullName": user.FullName,
		"email":    user.Email,
	})
}
