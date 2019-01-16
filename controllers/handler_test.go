package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bayugyug/sample-rates/config"
	"github.com/bayugyug/sample-rates/models"
	"github.com/bayugyug/sample-rates/utils"
)

var thandler *AppHandler

//TestHandler default initializer
func TestHandler(t *testing.T) {
	var err error
	t.Log("Init test")
	//init
	thandler = &AppHandler{}
	config.Settings = config.NewAppSettings()
	AppService, err = NewService(
		WithSvcOptAddress(":"+config.Settings.Config.HttpPort),
		WithSvcOptDbConf(&config.Settings.Config.Driver),
	)
	//check
	if err != nil {
		t.Fatal("Config might be missing", err)
	}
}

//TestHandlerLatest /rates/latest
func TestHandlerLatest(t *testing.T) {
	//setup
	ts := httptest.NewServer(http.HandlerFunc(thandler.LatestRatesHandler))
	defer ts.Close()

	mockLists := []struct {
		Method string
		URL    string
		Ctx    context.Context
		Body   io.Reader
	}{
		{
			Method: "GET",
			URL:    "/rates/latest",
			Body:   bytes.NewBufferString(``),
		},
	}
	var rates models.RatesLatestResponse
	for _, rec := range mockLists {
		ret, body := testRequest(t, ts, rec.Method, rec.URL, rec.Body, "")
		if ret.StatusCode != http.StatusOK {
			t.Fatalf("Request status:%d", ret.StatusCode)
		}
		if err := json.Unmarshal([]byte(body), &rates); err != nil {
			t.Fatalf("Response failed")
		}
		utils.Dumper(rates)
	}
}

//TestHandlerLatestWithDate /rates/yyyy-mm-dd
func TestHandlerLatestWithDate(t *testing.T) {
	//setup
	ts := httptest.NewServer(http.HandlerFunc(thandler.RatesHandler))
	defer ts.Close()

	mockLists := []struct {
		Method string
		URL    string
		Ctx    context.Context
		Body   io.Reader
	}{
		{
			Method: "GET",
			URL:    "/rates/2019-01-14",
			Body:   bytes.NewBufferString(``),
		},
	}
	var rates models.RatesLatestResponse
	for _, rec := range mockLists {
		ret, body := testRequest(t, ts, rec.Method, rec.URL, rec.Body, "")
		if ret.StatusCode != http.StatusOK {
			t.Fatalf("Request status:%d", ret.StatusCode)
		}
		if err := json.Unmarshal([]byte(body), &rates); err != nil {
			t.Fatalf("Response failed")
		}
		utils.Dumper(rates)
	}
	t.Log("OK")
}

//TestHandlerLatestWithDateInvalid /rates/yyyy-mm-dd
func TestHandlerLatestWithDateInvalid(t *testing.T) {
	//setup
	ts := httptest.NewServer(http.HandlerFunc(thandler.RatesHandler))
	defer ts.Close()

	mockLists := []struct {
		Method string
		URL    string
		Ctx    context.Context
		Body   io.Reader
	}{
		{
			Method: "GET",
			URL:    "/rates/2019-01-14x",
			Body:   bytes.NewBufferString(``),
		},
	}
	var response AppResponse
	for _, rec := range mockLists {
		ret, body := testRequest(t, ts, rec.Method, rec.URL, rec.Body, "")
		if ret.StatusCode != http.StatusOK {
			t.Fatalf("Request status:%d", ret.StatusCode)
		}
		if err := json.Unmarshal([]byte(body), &response); err != nil {
			t.Fatalf("Response failed")
		}
		utils.Dumper(response)
	}
}

//TestHandlerLatest /rates/analyze
func TestHandlerAnalyze(t *testing.T) {
	//setup
	ts := httptest.NewServer(http.HandlerFunc(thandler.AnalyzeRatesHandler))
	defer ts.Close()

	mockLists := []struct {
		Method string
		URL    string
		Ctx    context.Context
		Body   io.Reader
	}{
		{
			Method: "GET",
			URL:    "/rates/analyze",
			Body:   bytes.NewBufferString(``),
		},
	}
	var rates models.RatesAnalyzeResponse
	for _, rec := range mockLists {
		ret, body := testRequest(t, ts, rec.Method, rec.URL, rec.Body, "")
		if ret.StatusCode != http.StatusOK {
			t.Fatalf("Request status:%d", ret.StatusCode)
		}
		if err := json.Unmarshal([]byte(body), &rates); err != nil {
			t.Fatalf("Response failed")
		}
		utils.Dumper(rates)
	}
}

//testRequest test for http req
func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader, auth string) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	if auth != "" {
		req.Header.Add("Authorization", "Bearer "+auth)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	defer resp.Body.Close()

	return resp, string(respBody)
}
