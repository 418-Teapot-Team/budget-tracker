package app

import "github.com/gin-gonic/gin"

type errorResponse struct {
	Message string `json:"message"`
}

func (app *App) newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}
