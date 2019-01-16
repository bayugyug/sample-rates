package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/bayugyug/sample-rates/config"
	"github.com/bayugyug/sample-rates/controllers"
)

const (
	VERSION_MAJOR = "0.1"
	VERSION_MINOR = "0"
)

// Versioning detail information, set during build phase.
var (
	BuildTime  string
	AppVersion string
)

//internal system initialize
func init() {

	//uniqueness
	rand.Seed(time.Now().UnixNano())
	AppVersion = "Ver: " + VERSION_MAJOR + "." + VERSION_MINOR + "-" + BuildTime

}

func main() {

	start := time.Now()
	log.Println(AppVersion)

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
