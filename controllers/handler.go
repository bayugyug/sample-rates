package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bayugyug/sample-rates/config"
	"github.com/bayugyug/sample-rates/utils"
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
	app.ReplyErrContent(w, r, http.StatusOK, "Welcome!")
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
	//per date
	utils.Dumper("tran-date:" + tdate)
	app.LatestRates(w, r, tdate)
}

func (app *AppHandler) LatestRatesHandler(w http.ResponseWriter, r *http.Request) {
	//method:GET
	if !app.MethodAllowed(w, r, "GET") {
		return
	}
	//check it
	qpath := r.URL.Path[len("/rates/"):]
	utils.Dumper("latest:" + qpath)
	app.LatestRates(w, r, "")
}

func (app *AppHandler) AnalyzeRatesHandler(w http.ResponseWriter, r *http.Request) {
	//method:GET
	if !app.MethodAllowed(w, r, "GET") {
		return
	}
	//check it
	qpath := r.URL.Path[len("/rates/"):]
	app.ReplyErrContent(w, r, http.StatusOK, "analyze:"+qpath)
}

func (app *AppHandler) Handler404(w http.ResponseWriter, r *http.Request) {
	//404
	app.ReplyErrContent(w, r, http.StatusNotFound, "Invalid Endpoint")
}

//ReplyErrContent send 204 msg
//
//  http.StatusNoContent
//  http.StatusText(http.StatusNoContent)
func (app *AppHandler) ReplyErrContent(w http.ResponseWriter, r *http.Request, code int, msg string) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AppResponse{
		Code:   code,
		Status: msg,
	})
}

//ReplyContent send 200 with data
func (app *AppHandler) ReplyContent(w http.ResponseWriter, r *http.Request, reply string) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(reply))
}

//MethodAllowed check method here
func (app *AppHandler) MethodAllowed(w http.ResponseWriter, r *http.Request, method string) bool {
	//StatusMethodNotAllowed = 405
	switch r.Method {
	case http.MethodGet:
		if !strings.EqualFold(http.MethodGet, method) {
			app.ReplyErrContent(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
			return false
		}
		return true
	case http.MethodPost:
		if !strings.EqualFold(http.MethodPost, method) {
			app.ReplyErrContent(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
			return false
		}
		return true
	case http.MethodPut:
		if !strings.EqualFold(http.MethodPut, method) {
			app.ReplyErrContent(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
			return false
		}
		return true
	case http.MethodDelete:
		if !strings.EqualFold(http.MethodDelete, method) {
			app.ReplyErrContent(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
			return false
		}
		return true
	default:
		app.ReplyErrContent(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return false

	}
}
