package model

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}
