package model

type Product struct {
	Id       int    `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Color    string `json:"color"`
	Size     int    `json:"size"`
}
type Migrate struct {
}
