package model

type Response struct {
	Service string `json:"service"`
	Status  string  `json:"status"`
	Code 	int 	`json:"code"`
}