package dto

type OrderCreateDTO struct {
	UserID         uint64 `json:"user_id" binding:"required"`
	OrderDate      int    `json:"order_date" binding:"required"`
	Total          int    `json:"total" binding:"required"`
	PaymentMethod  string `json:"payment_method" binding:"required"`
	ShippingMethod string `json:"shipping_method" binding:"required"`
	AddressID      uint64 `json:"address_id" binding:"required"`
	Status         string `json:"status" binding:"required"`
}

type OrderUpdateDTO struct {
	ID             uint64 `json:"id" form:"id"`
	UserID         uint64 `json:"user_id"`
	OrderDate      int    `json:"order_date"`
	Total          int    `json:"total"`
	PaymentMethod  string `json:"payment_method"`
	ShippingMethod string `json:"shipping_method"`
	AddressID      uint64 `json:"address_id"`
	Status         string `json:"status"`
}
