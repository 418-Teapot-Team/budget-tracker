package app

import (
	budget "budget-tracker"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"sort"
	"strconv"
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
	//if budgetType == "" {
	//	app.newErrorResponse(c, http.StatusBadRequest, "please indicate the type of budget")
	//	return
	//}

	if budgetType != "income" && budgetType != "expenses" && budgetType != "all" && budgetType != "" {
		app.newErrorResponse(c, http.StatusNotFound, "possible types: income, expenses, all(\"\")")
		return
	} else if budgetType == "all" {
		budgetType = ""
	}

	takeAmountParam := c.Query("take")
	takeAmount := 0

	if takeAmountParam != "" {
		var err error
		takeAmount, err = strconv.Atoi(takeAmountParam)
		if err != nil {
			takeAmount = 0
		}
	}

	skipAmountParam := c.Query("skip")
	skipAmount := 0

	if takeAmountParam != "" {
		var err error
		skipAmount, err = strconv.Atoi(skipAmountParam)
		if err != nil {
			skipAmount = 0
		}
	}

	userId, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list, _ := app.s.ListsService.GetList(userId, budgetType, orderBy, sortedBy, takeAmount, skipAmount)

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
		sort.Ints(categoriesToSkip)

		for i, curr := 0, 0; i < 4-lenList; i++ {
			// skipping present categories
			for c := 0; c < len(categoriesToSkip); c++ {
				if curr+1 == categoriesToSkip[c] {
					curr++
					if c+1 == len(categoriesToSkip) || curr+1 != categoriesToSkip[c+1] {
						break
					}
				}
			}

			list = append(list, budget.ListsGetter{Categories: categories[curr]})
			curr++
		}
	}

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

func (app *App) getTotalAmount(c *gin.Context) {

	userId, err := app.getUserId(c)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	monthsParam := c.Query("months")
	months := 1

	if monthsParam != "" {
		var err error
		months, err = strconv.Atoi(monthsParam)
		if err != nil {
			months = 1
		}
	}

	budgetType := c.Query("type")
	if budgetType != "income" && budgetType != "expenses" {
		app.newErrorResponse(c, http.StatusNotFound, "Wrong type data. possible types: income, expenses")
		return
	}

	result, err := app.s.GetTotalAmount(userId, budgetType, months)
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": math.Round(result*100) / 100,
	})
}
