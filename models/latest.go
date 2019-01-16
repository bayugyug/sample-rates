package models

import (
	"context"
	"database/sql"
	"log"
)

type RateLatest struct {
}

func NewRateLatest() *RateLatest {
	return &RateLatest{}
}

func (rate *RateLatest) Get(ctx context.Context, db *sql.DB, whichdt string) *RatesLatestResponse {

	var rates *RatesLatestResponse
	var rows *sql.Rows
	var err error
	var r string

	//run
	if whichdt == "" {
		r = `SELECT 
				IFNULL(a.base,''),
				IFNULL(a.currency,''),
				IFNULL(a.rate,0.0)
			FROM rates a
			INNER JOIN (
			  SELECT base, MAX(rate_dt) AS rate_dt
			  FROM rates GROUP BY base
			) AS max_rate_dt
			USING (base, rate_dt)
			ORDER BY currency ASC`
		rows, err = db.QueryContext(ctx, r)
	} else {
		r = `SELECT 
		IFNULL(a.base,''), 
		IFNULL(a.currency,''), 
		IFNULL(a.rate,0.0)
		FROM rates a
		WHERE 1=1
		AND a.rate_dt = ?
		ORDER BY a.currency ASC`
		rows, err = db.QueryContext(ctx, r, whichdt)
	}

	if err != nil {
		log.Println("SQL_ERROR::", err)
		return rates
	}
	defer rows.Close()
	rates = &RatesLatestResponse{
		Rates: make(map[string]float64),
		Base:  "EUR",
	}
	for rows.Next() {
		var base, currency string
		var rate float64
		if err := rows.Scan(&base, &currency, &rate); err != nil {
			log.Println("SQL_ERROR::SCAN::", err)
			continue
		}
		//save it
		rates.Base = base
		rates.Rates[currency] = rate
	}
	if err := rows.Err(); err != nil {
		log.Println("SQL_ERROR::", err)
		return rates
	}
	//sounds good ;-)
	return rates
}
