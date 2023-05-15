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

func (app *App) authMiddleware(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		app.newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headersParts := strings.Split(header, " ")
	if len(headersParts) != 2 || headersParts[0] != "Bearer" {
		app.newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
	}

	if len(headersParts[1]) == 0 {
		app.newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := app.s.ParseToken(headersParts[1])
	if err != nil {
		app.newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	c.Set(userCtx, userId.UserId)
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

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		c.Next()
	}
}
