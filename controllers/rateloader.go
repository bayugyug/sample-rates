package controllers

import (
	"github.com/bayugyug/sample-rates/config"
	"github.com/bayugyug/sample-rates/models"
)

//LoadRates grab rates from remote
func (app *AppHandler) LoadRates() int {
	loader := models.NewRateLoader()
	results := loader.Load(AppService.Context, AppService.DB, config.Settings.RateEndPoints, config.Settings.Curl)
	var tot int
	for _, rate := range results {
		for _, one := range rate.Rates {
			if id := loader.Save(AppService.Context,
				AppService.DB,
				rate.Base,
				rate.Date,
				one.Currency,
				one.Rate); id > 0 {
				tot++
			}
		}
	}
	return tot
}
