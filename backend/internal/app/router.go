package app

import (
	"github.com/gin-gonic/gin"
)

func Routers(app *App) *gin.Engine {
	router := gin.New()

	router.Use(CORSMiddleware())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", app.signUp)
		auth.POST("/sign-in", app.signIn)
	}
	api := router.Group("/api", app.authMiddleware)
	{

		api.GET("/get-courses", app.getCourses)
		api.GET("/who-am-i", app.whoAmI)
		api.GET("/get-categories", app.getAllCategories)

		lists := api.Group("/lists")
		{
			lists.GET("/:type", app.getBudgetList)
			lists.POST("/", app.createList)
			lists.DELETE("/", app.deleteList)
			lists.GET("/get-top-expenses", app.getTopExpenses)
			lists.PUT("/", app.editBudgetList)
			lists.GET("/current-mon-saved", app.getCurrentMonthSavings)
		}

		accounts := api.Group("/accounts")
		{
			accounts.GET("/:account", app.getAllAccounts)
			accounts.POST("/", app.createAccount)
			accounts.DELETE("/", app.deleteAccount)
			accounts.PUT("/", app.editAccount)
		}
	}
	return router

}
