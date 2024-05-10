package model

import (
	"gorm.io/gorm"
)

type SubCategory struct {
	gorm.Model
	Name       string `gorm:"type:varchar(255)" json:"name"`
    CategoryID uint   `json:"category_id"`
}
