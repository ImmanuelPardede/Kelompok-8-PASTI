package dto

type AddressCreateDTO struct {
	UserID     uint   `json:"user_id" form:"user_id" binding:"required"`
	Street     string `json:"street" form:"street" binding:"required,min=3,max=255"`
	Village    string `json:"village" form:"village" binding:"required,min=3,max=100"`
	District   string `json:"district" form:"district" binding:"required,min=3,max=100"`
	Regency    string `json:"regency" form:"regency" binding:"required,min=3,max=100"`
	Province   string `json:"province" form:"province" binding:"required,min=3,max=100"`
	PostalCode int    `json:"postal_code" form:"postal_code" binding:"required"`
	Detail     string `json:"detail" form:"detail" binding:"required,min=3,max=255"`
}

type AddressUpdateDTO struct {
	ID         uint   `json:"id" form:"id" binding:"required"`
	UserID     uint   `json:"user_id" form:"user_id" binding:"required"`
	Street     string `json:"street" form:"street" binding:"required,min=3,max=255"`
	Village    string `json:"village" form:"village" binding:"required,min=3,max=100"`
	District   string `json:"district" form:"district" binding:"required,min=3,max=100"`
	Regency    string `json:"regency" form:"regency" binding:"required,min=3,max=100"`
	Province   string `json:"province" form:"province" binding:"required,min=3,max=100"`
	PostalCode int    `json:"postal_code" form:"postal_code" binding:"required"`
	Detail     string `json:"detail" form:"detail" binding:"required,min=3,max=255"`
}
