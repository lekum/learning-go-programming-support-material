package main

type Currency struct {
	Code    string `json:"currency_code"`
	Name    string `json:"currency_name"`
	Number  string `json:"currency_number"`
	Country string `json:"currency_country"`
}

type CurrencyRequest struct {
	Get   string `json:"get"`
	Limit int    `json:"limit"`
}
