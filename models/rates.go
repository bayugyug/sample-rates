package models

//for the responses

//RatesLatestResponse for /latest
type RatesLatestResponse struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

//RatesAnalyzeResponse for /analyze
type RatesAnalyzeResponse struct {
	Base         string             `json:"base"`
	RatesAnalyze map[string]RateAve `json:"rates_analyze"`
}

//RateAve map of min/max/ave
type RateAve struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Avg float64 `json:"avg"`
}

//for the xml parsers

//RateCurrency
type RateCurrency struct {
	Currency string  `json:"currency,omitempty"`
	Rate     float64 `json:"rate,omitempty"`
}

//RateRaw
type RateRaw struct {
	Base  string         `json:"base,omitempty"`
	Date  string         `json:"date,omitempty"`
	Rates []RateCurrency `json:"rates,omitempty"`
}
