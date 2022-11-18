package models

type Movement struct {
	ID       string `json:"ID"`
	Date     string `json:"date"`
	Type     string `json:"type"`
	Commerce string `json:"commerce"`
	Category string `json:"category"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}
