package model

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255)" json:"name"`
	Image string `gorm:"type:varchar(255)" json:"image"`
}
