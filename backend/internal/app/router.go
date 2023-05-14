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
	api := router.Group("/api", app.authMiddleware)
	{
		api.GET("/getAll", app.getCourses)

		api.GET("who-am-i", app.whoAmI)

	}


	}
	return router

}
