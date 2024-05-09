package repository

import (
	"github.com/iqbalsiagian17/Service_Brand/model"
	"gorm.io/gorm"
)

type BrandRepository interface {
	InsertBrand(brand model.Brand) model.Brand
	UpdateBrand(brand model.Brand) model.Brand
	All() []model.Brand
	FindByID(BrandID uint) model.Brand
	DeleteBrand(brand model.Brand)
}

type brandConnection struct {
	connection *gorm.DB
}

func NewBrandRepository(db *gorm.DB) BrandRepository {
	return &brandConnection{
		connection: db,
	}
}

func (db *brandConnection) InsertBrand(brand model.Brand) model.Brand {
	db.connection.Save(&brand)
	return brand
}

func (db *brandConnection) UpdateBrand(brand model.Brand) model.Brand {
	db.connection.Save(&brand)
	return brand
}

func (db *brandConnection) All() []model.Brand {
	var brands []model.Brand
	db.connection.Find(&brands)
	return brands
}

func (db *brandConnection) FindByID(brandID uint) model.Brand {
	var brand model.Brand
	db.connection.Find(&brand, brandID)
	return brand
}

func (db *brandConnection) DeleteBrand(brand model.Brand) {
	db.connection.Delete(&brand)
}
