package model

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	UserID     uint   `json:"user_id"`
	Street     string `gorm:"type:varchar(255)" json:"street"`
	Village    string `gorm:"type:varchar(100)" json:"village"`
	District   string `gorm:"type:varchar(100)" json:"district"`
	Regency    string `gorm:"type:varchar(100)" json:"regency"`
	Province   string `gorm:"type:varchar(100)" json:"province"`
	PostalCode int    `json:"postal_code"`
	Detail     string `gorm:"type:varchar(255)" json:"detail"`
}
