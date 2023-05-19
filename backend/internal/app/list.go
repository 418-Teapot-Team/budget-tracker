package app

import (
	budget "budget-tracker"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

func (app *App) createList(c *gin.Context) {
	var input budget.List

	userId, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = c.BindJSON(&input); err != nil {
		app.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err = app.s.ListsService.CreateList(&input, userId)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

type deleteInput struct {
	Id int `json:"id"`
}

func (app *App) deleteList(c *gin.Context) {
	var input deleteInput

	userId, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = c.BindJSON(&input); err != nil {
		app.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = app.s.ListsService.DeleteList(input.Id, userId)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})

}

func (app *App) getBudgetList(c *gin.Context) {

	orderBy := c.Query("orderBy")
	if orderBy == "" {
		orderBy = "created_at"
	}

	sortedBy := c.Query("reverse")
	if sortedBy == "" || sortedBy == "false" {
		sortedBy = " desc"
	} else if sortedBy == "true" {
		sortedBy = " asc"
	} else {
		app.newErrorResponse(c, http.StatusBadRequest, "invalid query param reverse")
		return
	}

	budgetType := c.Param("type")
	if budgetType == "" {
		app.newErrorResponse(c, http.StatusBadRequest, "please indicate the type of budget")
		return
	}

	if budgetType != "income" && budgetType != "expenses" {
		budgetType = ""
	}
	userId, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list, _ := app.s.ListsService.GetList(userId, budgetType, orderBy, sortedBy)

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": list,
	})
}

func (app *App) editBudgetList(c *gin.Context) {
	var input budget.List

	userid, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = c.BindJSON(&input); err != nil {
		app.newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.UserId = userid

	err = app.s.ListsService.EditList(input)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, input)

}

func (app *App) getTopCategories(c *gin.Context) {

	listType := c.Query("type")
	if listType != "income" && listType != "expenses" {
		app.newErrorResponse(c, http.StatusInternalServerError, "define a 'type' query value (income or expenses)")
		return

	}

	userId, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list, err := app.s.ListsService.GetTopListsCategories(userId, listType)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if listType == "expenses" && len(list) < 4 {
		lenList := len(list)
		categories, err := app.s.CategoriesService.GetAllCategories()
		if err != nil {
			app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		var categoriesToSkip []int
		for i := 0; i < lenList; i++ {
			categoriesToSkip = append(categoriesToSkip, list[i].Category)
		}

		for i, curr := 0, 0; i < 4-lenList; i++ {
			for _, j := range categoriesToSkip {
				// skipping category if already present
				if j == curr {
					curr++
				}
			}
			list = append(list, budget.ListsGetter{Categories: categories[curr]})
			curr++
		}
	}

	fmt.Println(list)

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": list,
	})
}

func (app *App) getCurrentMonthSavings(c *gin.Context) {

	userId, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	incomeNet, err := app.s.GetCurrentMonthSavings(userId)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": math.Round(incomeNet*100) / 100,
	})
}

func (app *App) getSavingsStats(c *gin.Context) {

	userId, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	data, err := app.s.GetSavingsStats(userId)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println(data)

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": data,
	})
}
