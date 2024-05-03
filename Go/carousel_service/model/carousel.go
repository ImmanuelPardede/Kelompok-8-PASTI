package model

import (
	"gorm.io/gorm"
)

type Carousel struct {
	gorm.Model
	ID           uint64 `gorm:"primary_key:auto_increment"`
	ImageURL   string `gorm:"column:image_url"`
	Caption    string `gorm:"column:caption"`
	Subcaption string `gorm:"column:subcaption"`
}
