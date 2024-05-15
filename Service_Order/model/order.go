package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID             uint64 `gorm:"primaryKey" json:"id"`
	UserID         uint64 `gorm:"type:int" json:"user_id"`
	OrderDate      int    `gorm:"type:int" json:"order_date"`
	Total          int    `gorm:"type:int" json:"total"`
	PaymentMethod  string `gorm:"type:varchar(255)" json:"payment_method"`
	ShippingMethod string `gorm:"type:varchar(50)" json:"shipping_method"`
	AddressID      uint64   `gorm:"type:int" json:"shipping_address"`
	Status         string `gorm:"type:varchar(50)" json:"status"`
}
