package model

type Response struct {
	Service string `json:"service"`
	Status  bool   `json:"status"`
}