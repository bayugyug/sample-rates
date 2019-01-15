package models

type Rate struct {
	Base  string             `json:"base,omitempty"`
	Rates map[string]float64 `json:"rates,omitempty"`
}

type RateInfo struct {
	Base         string             `json:"base,omitempty"`
	RatesAnalyze map[string]RateAve `json:"rates_analyze,omitempty"`
}

type RateAve struct {
	Min float64 `json:"min,omitempty"`
	Max float64 `json:"max,omitempty"`
	Avg float64 `json:"avg,omitempty"`
}


type RateCurrency struct {
	Currency string  `json:"currency,omitempty"`
	Rate     float64 `json:"rate,omitempty"`
}

type RateRaw struct {
	Base  string         `json:"base,omitempty"`
	Date  string         `json:"date,omitempty"`
	Rates []RateCurrency `json:"rates,omitempty"`
}
