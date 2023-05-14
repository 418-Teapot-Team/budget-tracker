package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (app *App) Middleware(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		app.newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headersParts := strings.Split(header, "")
	if len(headersParts) != 2 || headersParts[0] != "Bearer" {
		app.newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
	}

	if len(headersParts[1]) == 0 {
		app.newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := app.s.ParseToken(headersParts[1])
	if err != nil {
		app.newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	c.Set(userCtx, userId)
}

func (app *App) getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	integer, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return integer, nil
}
