package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (app *App) getCourses(c *gin.Context) {

	currencies := []string{"USD", "EUR", "PLN"}

	var rates []courses

	off, err := app.client.GetOfficialRate()
	if err != nil {
		log.Println(err)
		app.newErrorResponse(c, http.StatusInternalServerError, "error while trying to get data from nbu.ua")
		return
	}
	dark, err := app.client.GetGoverlaRate()
	if err != nil {
		app.newErrorResponse(c, http.StatusInternalServerError, "error while trying to get data from goverla.ua")
		return
	}

	for _, rt := range dark.Data.Point.Rates {
		if contains(rt.Currency.CodeAlpha, currencies) {
			var (
				ofRate string
			)
			for _, of := range off {
				if of.Cc == rt.Currency.CodeAlpha {
					ofRate = fmt.Sprint(of.Rate)
				}
			}
			rate := fmt.Sprint(rt.Ask.Absolute)

			darkRate := rate[:len(rate)-rt.Currency.Exponent] + "." + rate[rt.Currency.Exponent:]
			rates = append(rates,
				courses{
					Code:         rt.Currency.CodeAlpha,
					DarkRate:     darkRate,
					OfficialRate: ofRate,
				})
		}
	}

	c.JSON(http.StatusOK, rates)

}

type courses struct {
	Txt          string `json:"txt"`
	Code         string `json:"code"`
	OfficialRate string `json:"officialRate"`
	DarkRate     string `json:"darkRate"`
}

func contains(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false

}
