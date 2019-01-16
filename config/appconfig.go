package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"time"

	"github.com/bayugyug/sample-rates/driver"
	"github.com/bayugyug/sample-rates/utils"
)

// Versioning detail information, set during build phase.
var (
	Settings *AppSettings
)

//internal system initialize
func init() {

	//uniqueness
	rand.Seed(time.Now().UnixNano())

}

const (
	//status
	usageConfig = "use to set the config file parameter with db-userinfos/http-port"
)

//RateEndPoint rate url mapping
type RateEndPoint struct {
	Base string
	URL  string
}

//ParameterConfig optional parameter structure
type ParameterConfig struct {
	Driver   driver.DbConnectorConfig `json:"driver"`
	HttpPort string                   `json:"http_port"`
	Showlog  bool                     `json:"showlog"`
}

//AppSettings app mapping on its config
type AppSettings struct {
	Config        *ParameterConfig
	CmdParams     string
	EnvVars       map[string]*string
	RateEndPoints []RateEndPoint
	Curl          *utils.HttpCurl
	DateRegx      *regexp.Regexp
}

//NewAppSettings main entry for config
func NewAppSettings() *AppSettings {
	//set default
	cfg := &AppSettings{
		EnvVars: make(map[string]*string),
	}
	//maybe export from envt
	cfg.EnvVars = map[string]*string{
		"SAMPLE_RATES_CONFIG": &cfg.CmdParams,
	}
	//start
	cfg.Initializer()
	return cfg
}

//InitRecov is for dumpIng segv in
func (g *AppSettings) InitRecov() {
	//might help u
	defer func() {
		recvr := recover()
		if recvr != nil {
			fmt.Println("MAIN-RECOV-INIT: ", recvr)
		}
	}()
}

//InitEnvParams enable all OS envt vars to reload internally
func (g *AppSettings) InitEnvParams() {
	//just in-case, over-write from ENV
	for k, v := range g.EnvVars {
		if os.Getenv(k) != "" {
			*v = os.Getenv(k)
		}
	}
	//get options
	flag.StringVar(&g.CmdParams, "config", g.CmdParams, usageConfig)
	flag.Parse()
}

//Initializer set defaults for initial reqmts
func (g *AppSettings) Initializer() {
	//prepare
	g.InitRecov()
	g.InitEnvParams()

	//try to reconfigure if there is passed params, otherwise use the default
	if g.CmdParams != "" {
		g.Config = g.FormatParameterConfig(g.CmdParams)
	}

	//load defaults
	if g.Config == nil {
		g.Config = &ParameterConfig{
			Driver: driver.DbConnectorConfig{
				User: "rates",
				Pass: "rat3s",
				Host: "127.0.0.1",
				Port: "3306",
				Name: "rates",
			},
			HttpPort: "8989",
			Showlog:  true,
		}
	}

	//set dump flag
	utils.ShowMeLog = g.Config.Showlog

	//set endpoints
	g.RateEndPoints = []RateEndPoint{
		{
			Base: "EUR",
			URL:  "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml",
		},
	}
	//http
	g.Curl = utils.NewHttpCurl()
	//date grep
	g.DateRegx = regexp.MustCompile("^[0-9]{4}-[0-9]{2}-[0-9]{2}$")
}

//FormatParameterConfig new ParameterConfig
func (g *AppSettings) FormatParameterConfig(s string) *ParameterConfig {
	var cfg ParameterConfig
	if err := json.Unmarshal([]byte(s), &cfg); err != nil {
		log.Println("FormatParameterConfig", err)
		return nil
	}
	return &cfg
}
