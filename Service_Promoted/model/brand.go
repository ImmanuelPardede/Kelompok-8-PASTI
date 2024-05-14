package model

import (
	"gorm.io/gorm"
)

type Promoted struct {
	gorm.Model
	Title string `gorm:"type:varchar(255)" json:"title"`
	Image string `gorm:"type:varchar(255)" json:"image"`
}
