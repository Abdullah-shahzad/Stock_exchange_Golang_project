package models

type Transaction struct {
	ID               int     `json:"id"`
	UserID           int     `json:"user_id"`
	Ticker           string  `json:"ticker"`
	TransactionType  string  `json:"transaction_type"`
	TransactionVolume int    `json:"transaction_volume"`
	TransactionPrice  float64 `json:"transaction_price"`
	Timestamp         string  `json:"timestamp"`
}
