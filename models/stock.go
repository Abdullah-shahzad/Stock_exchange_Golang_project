package models

type Stock struct {
	ID    int     `json:"id"`
	Ticker string `json:"ticker"`
	Price  float64 `json:"price"`
}
