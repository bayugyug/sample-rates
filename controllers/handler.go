package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bayugyug/sample-rates/config"
)

const (
	MsgStatusOK  = "Success"
	MsgStatusNOK = "Error"
)

type AppResponse struct {
	Code   int
	Status string
}

type AppHandler struct {
}

func (app *AppHandler) WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	//reply
	app.ReplyContent(w, r, http.StatusOK, "Welcome!")
}

func (app *AppHandler) RatesHandler(w http.ResponseWriter, r *http.Request) {
	//method:GET
	if !app.MethodAllowed(w, r, "GET") {
		return
	}
	//for now, just send 404
	if !strings.HasPrefix(r.URL.Path, "/rates/") {
		app.Handler404(w, r)
		return
	}
	//yyyy-mm-dd
	tdate := r.URL.Path[len("/rates/"):]
	if !config.Settings.DateRegx.MatchString(tdate) {
		app.Handler404(w, r)
		return
	}
	//	row, err := driver.GetDriver(ApiService.Context, ApiService.DB, data.MobileDriver)
	app.ReplyContent(w, r, http.StatusOK, "tran-date:"+tdate)
}

func (app *AppHandler) LatestRatesHandler(w http.ResponseWriter, r *http.Request) {
	//method:GET
	if !app.MethodAllowed(w, r, "GET") {
		return
	}
	//check it
	qpath := r.URL.Path[len("/rates/"):]

	app.ReplyContent(w, r, http.StatusOK, "latest:"+qpath)
}

func (app *AppHandler) AnalyzeRatesHandler(w http.ResponseWriter, r *http.Request) {
	//method:GET
	if !app.MethodAllowed(w, r, "GET") {
		return
	}
	//check it
	qpath := r.URL.Path[len("/rates/"):]

	app.ReplyContent(w, r, http.StatusOK, "analyze:"+qpath)
}

func (app *AppHandler) Handler404(w http.ResponseWriter, r *http.Request) {
	//404
	app.ReplyContent(w, r, http.StatusNotFound, "Invalid Endpoint")
}

//ReplyContent send 204 msg
//
//  http.StatusNoContent
//  http.StatusText(http.StatusNoContent)
func (app *AppHandler) ReplyContent(w http.ResponseWriter, r *http.Request, code int, msg string) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AppResponse{
		Code:   code,
		Status: msg,
	})
}

//MethodAllowed check method here
func (app *AppHandler) MethodAllowed(w http.ResponseWriter, r *http.Request, method string) bool {
	//StatusMethodNotAllowed = 405
	switch r.Method {
	case http.MethodGet:
		if !strings.EqualFold(http.MethodGet, method) {
			app.ReplyContent(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
			return false
		}
		return true
	case http.MethodPost:
		if !strings.EqualFold(http.MethodPost, method) {
			app.ReplyContent(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
			return false
		}
		return true
	case http.MethodPut:
		if !strings.EqualFold(http.MethodPut, method) {
			app.ReplyContent(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
			return false
		}
		return true
	case http.MethodDelete:
		if !strings.EqualFold(http.MethodDelete, method) {
			app.ReplyContent(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
			return false
		}
		return true
	default:
		app.ReplyContent(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return false

	}
}


