package repository

import (
	"github.com/iqbalsiagian17/Service_Product/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertProduct(product model.Product) model.Product
	UpdateProduct(product model.Product) model.Product
	All() []model.Product
	FindByID(ProductID uint) model.Product
	DeleteProduct(product model.Product)
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productConnection{
		connection: db,
	}
}

func (db *productConnection) InsertProduct(product model.Product) model.Product {
	db.connection.Save(&product)
	return product
}

func (db *productConnection) UpdateProduct(product model.Product) model.Product {
	db.connection.Save(&product)
	return product
}

func (db *productConnection) All() []model.Product {
	var products []model.Product
	db.connection.Find(&products)
	return products
}

func (db *productConnection) FindByID(productID uint) model.Product {
	var product model.Product
	db.connection.Find(&product, productID)
	return product
}

func (db *productConnection) DeleteProduct(product model.Product) {
	db.connection.Delete(&product)
}
