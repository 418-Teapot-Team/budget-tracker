package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunApp(app *App, port string) {
	router := gin.Default()

	router.Use(cors.Default())

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
			lists.POST("/create", app.createList)
			lists.DELETE("/delete", app.deleteList)
			lists.PUT("/update", app.editBudgetList)

			lists.GET("/get-top-categories", app.getTopCategories)
			lists.GET("/current-mon-saved", app.getCurrentMonthSavings)
			lists.GET("/get-saving-stats", app.getSavingsStats)
			lists.GET("/get-total-amount", app.getTotalAmount)
			lists.GET("/get-stats", app.getStats)
		}

		accounts := api.Group("/accounts")
		{
			accounts.GET("/:account", app.getAllAccounts)
			accounts.POST("/create", app.createAccount)
			accounts.DELETE("/delete", app.deleteAccount)
			accounts.PUT("/update", app.editAccount)

			accounts.GET("/total", app.getTotalDeposits)
		}
	}
	router.Run(":" + port)

}
