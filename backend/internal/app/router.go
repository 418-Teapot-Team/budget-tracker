package app

import (
	"github.com/gin-gonic/gin"
)

func Routers(app *App) *gin.Engine {
	router := gin.New()

	router.Use(CorsMiddleware())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", app.signUp)
		auth.POST("/sign-in", app.signIn)
	}
	api := router.Group("/api", app.authMiddleware)
	{
		api.GET("/get-courses", app.getCourses)
		api.GET("/who-am-i", app.whoAmI)
		lists := api.Group("/lists")
		{
			lists.GET("/:type", app.getBudgetList)
			lists.POST("/", app.createList)
			lists.DELETE("/", app.deleteList)
		}
	}
	return router

}
