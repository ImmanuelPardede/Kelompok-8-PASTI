package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string `gorm:"type:varchar(255)" json:"name"`
	CategoryID    uint   `gorm:"type:int" json:"category_id"`
	SubCategoryID uint   `gorm:"type:int" json:"subcategory_id"`
	BrandID       uint   `gorm:"type:int" json:"brand_id"`
	Size          string `gorm:"type:varchar(50)" json:"size"`
	Quantity      int    `gorm:"type:int" json:"quantity"`
	Price         int    `gorm:"type:int" json:"price"`
	Description   string `gorm:"type:text" json:"description"`
	Image         string `gorm:"type:varchar(255)" json:"image"`
}
