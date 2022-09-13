package model

import (
	"time"
)

type Product struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Price     int       `json:"price"`
	Color     string    `json:"color"`
	Size      int       `json:"size"`
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time
}
