package model

type Product struct {
	Id       int    `json:"id"`
	code     string `json:"code"`
	name     string `json:"name"`
	category string `json:"category"`
	Price    int    `json:"price"`
}
