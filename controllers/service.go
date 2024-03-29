package controllers

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bayugyug/sample-rates/config"
	"github.com/bayugyug/sample-rates/driver"
)

const (
	svcOptionWithAddress  = "svc-opts-address"
	svcOptionWithDbConfig = "svc-opts-db-config"
	svcOptionWithHandler  = "svc-opts-handler"
)

type Service struct {
	Address string
	App     *AppHandler
	DB      *sql.DB
	Config  *driver.DbConnectorConfig
	Context context.Context
}

//api global handler
var AppService *Service

//WithSvcOptAddress opts for port#
func WithSvcOptAddress(r string) *config.Option {
	return config.NewOption(svcOptionWithAddress, r)
}

//WithSvcOptDbConf opts for db connector
func WithSvcOptDbConf(r *driver.DbConnectorConfig) *config.Option {
	return config.NewOption(svcOptionWithDbConfig, r)
}

//WithSvcOptHandler opts for handler
func WithSvcOptHandler(r *AppHandler) *config.Option {
	return config.NewOption(svcOptionWithHandler, r)
}

//NewService service new instance
func NewService(opts ...*config.Option) (*Service, error) {

	//default
	svc := &Service{
		Address: ":8989",
		App:     &AppHandler{},
		Context: context.Background(),
	}

	//add options if any
	for _, o := range opts {
		switch o.Name() {
		case svcOptionWithHandler:
			if s, ok := o.Value().(*AppHandler); ok && s != nil {
				svc.App = s
			}
		case svcOptionWithAddress:
			if s, ok := o.Value().(string); ok && s != "" {
				svc.Address = s
			}
		case svcOptionWithDbConfig:
			if s, ok := o.Value().(*driver.DbConnectorConfig); ok && s != nil {

				svc.Config = s
			}
		}
	}

	//get db
	dbh, err := driver.NewDbConnector(svc.Config)
	if err != nil {
		return svc, err
	}

	//save
	svc.DB = dbh

	return svc, nil
}

//PrepareRates try to load rates here
func (svc *Service) PrepareRates() {
	log.Println("Initializing rates")
	log.Println("Wait please ....")
	ret := svc.App.LoadRates()
	log.Println("Rates initialized#", ret)

}

//Run run the http server based on settings
func (svc *Service) Run() {

	//routing
	mux := http.NewServeMux()
	mux.HandleFunc("/rates/analyze", svc.App.AnalyzeRatesHandler)
	mux.HandleFunc("/rates/analyze/", svc.App.AnalyzeRatesHandler)
	mux.HandleFunc("/rates/latest", svc.App.LatestRatesHandler)
	mux.HandleFunc("/rates/latest/", svc.App.LatestRatesHandler)
	mux.HandleFunc("/rates/", svc.App.RatesHandler)
	mux.HandleFunc("/rates", svc.App.RatesHandler)
	mux.HandleFunc("/", svc.App.WelcomeHandler)

	//gracious timing
	srv := &http.Server{
		Addr:         svc.Address,
		Handler:      mux,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	//async run
	go func() {
		log.Println("Listening on port ", svc.Address)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
			os.Exit(0)
		}

	}()

	//watcher
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	<-stopChan
	log.Println("Shutting down service...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()
	log.Println("Server gracefully stopped!")

}

//SetContextKeyVal version context
func (svc *Service) SetContextKeyVal(k, v string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), k, v))
			next.ServeHTTP(w, r)
		})
	}
}
