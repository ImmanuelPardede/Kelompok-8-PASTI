package model

import ("gorm.io/gorm")

type Cart struct {
	gorm.Model
	ProductID string `json:"product_id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Quantity int64 `json:"quantity"`
	Items int64 `json:"items"`
}
