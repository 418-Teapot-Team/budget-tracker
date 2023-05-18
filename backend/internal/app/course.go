package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *App) getCourses(c *gin.Context) {

	currencies := []string{"USD", "EUR", "PLN"}

	var rates []courses

	dark, err := app.client.GetGoverlaRate()
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, "error while trying to get data from goverla.ua")
		return
	}

	for _, rt := range dark.Data.Point.Rates {
		if contains(rt.Currency.CodeAlpha, currencies) {

			buy := fmt.Sprint(rt.Ask.Absolute)
			sale := fmt.Sprint(rt.Bid.Absolute)

			buyRate := buy[:len(buy)-rt.Currency.Exponent] + "." + buy[len(buy)-rt.Currency.Exponent:]
			saleRate := sale[:len(sale)-rt.Currency.Exponent] + "." + sale[len(sale)-rt.Currency.Exponent:]
			rates = append(rates,
				courses{
					Code:     rt.Currency.CodeAlpha,
					BuyRate:  buyRate,
					SaleRate: saleRate,
				})
		}
	}

	c.JSON(http.StatusOK, rates)

}

type courses struct {
	Code     string `json:"code"`
	BuyRate  string `json:"buyRate"`
	SaleRate string `json:"saleRate"`
}

func contains(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false

}
