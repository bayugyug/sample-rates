package models

import (
	"context"
	"database/sql"
	"log"
)

type RateAnalyze struct {
}

func NewRateAnalyze() *RateAnalyze {
	return &RateAnalyze{}
}

func (rate *RateAnalyze) Get(ctx context.Context, db *sql.DB) *RatesAnalyzeResponse {

	var rates *RatesAnalyzeResponse
	var rows *sql.Rows
	var err error

	//run
	r := `SELECT
			IFNULL(a.base,'')       as base,
			IFNULL(a.currency,'')   as currency,
			MIN(IFNULL(a.rate,0.0)) as minimum,
			MAX(IFNULL(a.rate,0.0)) as maximum,
			AVG(IFNULL(a.rate,0.0)) as average
		FROM
			rates a
			WHERE 1=1
			GROUP BY base,currency ASC
		`
	rows, err = db.QueryContext(ctx, r)
	if err != nil {
		log.Println("SQL_ERROR::", err)
		return rates
	}
	defer rows.Close()
	rates = &RatesAnalyzeResponse{
		RatesAnalyze: make(map[string]RateAve),
		Base:         "EUR",
	}
	for rows.Next() {
		var base, currency string
		var ratemin, ratemax, rateavg float64
		if err := rows.Scan(&base, &currency, &ratemin, &ratemax, &rateavg); err != nil {
			log.Println("SQL_ERROR::SCAN::", err)
			continue
		}
		//save it
		rates.Base = base
		rates.RatesAnalyze[currency] = RateAve{
			Min: ratemin,
			Max: ratemax,
			Avg: rateavg,
		}
	}
	if err := rows.Err(); err != nil {
		log.Println("SQL_ERROR::", err)
		return rates
	}
	//sounds good ;-)
	return rates
}
