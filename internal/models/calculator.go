package models

type ConvertRequest struct {
    Number int `json:"number"`
}

type ConvertResponse struct {
    Decimal int    `json:"decimal"`
    Binary  string `json:"binary"`
    Error   string `json:"error,omitempty"`
}