package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *App) getAllCategories(c *gin.Context) {
	_, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	categories, err := app.s.CategoriesService.GetAllCategories()
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": categories,
	})
}
