package repository

import (
	"github.com/iqbalsiagian17/Service_Cart/model"
	"gorm.io/gorm"
)

type CartRepository interface {
	InsertCart(cart model.Cart) model.Cart
	UpdateCart(cart model.Cart) model.Cart
	All() []model.Cart
	FindByID(cartID uint) model.Cart
	DeleteCart(cart model.Cart)
}

type cartConnection struct {
	connection *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartConnection{
		connection: db,
	}
}

func (db *cartConnection) InsertCart(cart model.Cart) model.Cart {
	db.connection.Create(&cart)
	return cart
}

func (db *cartConnection) UpdateCart(cart model.Cart) model.Cart {
	db.connection.Save(&cart)
	return cart
}

func (db *cartConnection) All() []model.Cart {
	var carts []model.Cart
	db.connection.Find(&carts)
	return carts
}

func (db *cartConnection) FindByID(cartID uint) model.Cart {
	var cart model.Cart
	db.connection.First(&cart, cartID)
	return cart
}

func (db *cartConnection) DeleteCart(cart model.Cart) {
	db.connection.Delete(&cart)
}
