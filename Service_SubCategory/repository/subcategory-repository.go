package repository

import (
	"github.com/iqbalsiagian17/Service_SubCategory/model"
	"gorm.io/gorm"
)

type SubcategoryRepository interface {
	InsertSubcategory(subcategory model.SubCategory) model.SubCategory
	UpdateSubcategory(subcategory model.SubCategory) model.SubCategory
	All() []model.SubCategory
	FindByID(SubcategoryID uint) model.SubCategory
	DeleteSubcategory(subcategory model.SubCategory)
}

type subcategoryConnection struct {
	connection *gorm.DB
}

func NewSubCategoryRepository(db *gorm.DB) SubcategoryRepository {
	return &subcategoryConnection{
		connection: db,
	}
}

func (db *subcategoryConnection) InsertSubcategory(subcategory model.SubCategory) model.SubCategory {
	db.connection.Save(&subcategory)
	return subcategory
}

func (db *subcategoryConnection) UpdateSubcategory(subcategory model.SubCategory) model.SubCategory {
	db.connection.Save(&subcategory)
	return subcategory
}

func (db *subcategoryConnection) All() []model.SubCategory {
	var subcategories []model.SubCategory
	db.connection.Find(&subcategories)
	return subcategories
}

func (db *subcategoryConnection) FindByID(subcategoryID uint) model.SubCategory {
	var subcategory model.SubCategory
	db.connection.Find(&subcategory, subcategoryID)
	return subcategory
}

func (db *subcategoryConnection) DeleteSubcategory(subcategory model.SubCategory) {
	db.connection.Delete(&subcategory)
}
