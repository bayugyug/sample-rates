package models

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/bayugyug/sample-rates/config"
	"github.com/bayugyug/sample-rates/utils"
)

type RateLoader struct {
}

func NewRateLoader() *RateLoader {
	return &RateLoader{}
}

func (rate *RateLoader) Load(ctx context.Context, db *sql.DB, endpoints []config.RateEndPoint, curl *utils.HttpCurl) []RateRaw {
	var all []RateRaw
	for _, point := range endpoints {
		body, code, err := curl.Get(point.URL)
		if code != http.StatusOK || err != nil {
			log.Println("Oops, download xml failed!", code, err, point.URL)
			continue
		}
		v := XMLEnvelope{}
		if err := xml.Unmarshal([]byte(body), &v); err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		for _, cubeInfo1 := range v.Cubes {
			rateraw := RateRaw{
				Date: cubeInfo1.Time,
				Base: point.Base,
			}
			for _, cubeInfo2 := range cubeInfo1.Cubes {
				rt, _ := strconv.ParseFloat(cubeInfo2.Rate, 64)
				rateraw.Rates = append(rateraw.Rates, RateCurrency{
					Currency: cubeInfo2.Currency,
					Rate:     rt,
				})
			}
			//result
			all = append(all, rateraw)
		}
	}
	//good;-)
	return all
}


func (rate *RateLoader) Save(ctx context.Context, db *sql.DB, rbase,rdate, rcurrency string, rrate float64)int64 {
	//fmt
	r := `INSERT INTO rates (
		base, 
		rate_dt, 
		currency, 
		rate, 
		created_dt)
	      VALUES (?, ?, ?, ?,Now()) 
	      ON DUPLICATE KEY UPDATE
			base     = ? , 
			rate_dt  = ? , 
			currency = ? , 
			rate     = ? , 
			modified_dt = Now() `
	//exec
	result, err := db.Exec(r,
		rbase,
		rdate,
		rcurrency,
		rrate,
		rbase, // update starts here
		rdate,
		rcurrency,
		rrate,
	)
	if err != nil {
		log.Println("SQL_ERR", err)
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil || id < 1 {
		log.Println("SQL_ERR", err)
		return 0
	}
	//sounds good ;-)
	return int64(id)
}