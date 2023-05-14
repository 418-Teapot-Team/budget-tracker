package app

import (
	"github.com/gin-gonic/gin"
)

func Routers(app *App) *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", app.signUp)
		auth.POST("/sign-in", app.signIn)
	}
	rate := router.Group("/courses")
	{
		rate.GET("/getAll", app.getCourses)
	}
	return router

}
