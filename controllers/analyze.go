package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bayugyug/sample-rates/models"
)

//AnalyzeRates grab rates latest from db
func (app *AppHandler) AnalyzeRates(w http.ResponseWriter, r *http.Request) {
	latest := models.NewRateAnalyze()
	results := latest.Get(AppService.Context, AppService.DB)

	//NO_RECORDS
	if results == nil {
		//404
		app.ReplyErrContent(w, r, http.StatusNotFound, "No rates found")
		return
	}
	//good ;-)
	jdata, err := json.Marshal(results)
	if err != nil {
		app.ReplyErrContent(w, r, http.StatusInternalServerError, "Formatting failed")
		return
	}
	app.ReplyContent(w, r, string(jdata))
}
