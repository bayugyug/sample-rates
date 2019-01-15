package main

import (
	"log"
	"time"

	"github.com/bayugyug/sample-rates/config"
	"github.com/bayugyug/sample-rates/controllers"
)

func main() {

	start := time.Now()
	log.Println("Ver: ", config.AppVersion)

	var err error

	//init
	config.Settings = config.NewAppSettings()
	controllers.AppService, err = controllers.NewService(
		controllers.WithSvcOptAddress(":"+config.Settings.Config.HttpPort),
		controllers.WithSvcOptDbConf(&config.Settings.Config.Driver),
	)

	//check
	if err != nil {
		log.Fatal("Oops! config might be missing", err)
	}
	
	//prep-rates
	controllers.AppService.PrepareRates()
	
	//run service
	controllers.AppService.Run()
	log.Println("Since", time.Since(start))
	log.Println("Done")
}
