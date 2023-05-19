package app

import (
	"github.com/gin-gonic/gin"
)

func Routers(app *App) *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware())

	auth := router.Group("/auth")
	auth.Use(CORSMiddleware())
	{
		auth.POST("/sign-up", app.signUp)
		auth.POST("/sign-in", app.signIn)
	}
	api := router.Group("/api", app.authMiddleware)
	api.Use(CORSMiddleware())
	{

		api.GET("/get-courses", app.getCourses)
		api.GET("/who-am-i", app.whoAmI)
		api.GET("/get-categories", app.getAllCategories)

		lists := api.Group("/lists")
		{
			lists.GET("/:type", app.getBudgetList)
			lists.POST("/create", app.createList)
			lists.DELETE("/delete", app.deleteList)
			lists.PUT("/update", app.editBudgetList)

			lists.GET("/get-top-categories", app.getTopCategories)
			lists.GET("/current-mon-saved", app.getCurrentMonthSavings)
			lists.GET("/get-saving-stats", app.getSavingsStats)
			lists.GET("/get-total-amount", app.getTotalAmount)
			lists.GET("/get-stats", app.getStats)
		}

		api.GET("/account-total", app.getTotalDeposits)

		accounts := api.Group("/accounts")
		{

			accounts.GET("/:account", app.getAllAccounts)
			accounts.POST("/create", app.createAccount)
			accounts.DELETE("/delete", app.deleteAccount)
			accounts.PUT("/update", app.editAccount)
		}
	}
	return router
}
