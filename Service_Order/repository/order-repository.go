package repository

import (
	"github.com/iqbalsiagian17/Service_Order/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Insert(order model.Order) model.Order
	Update(order model.Order) model.Order
	Delete(order model.Order)
	All() []model.Order
	FindByID(orderID uint) model.Order
}

type orderConnection struct {
	connection *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderConnection{
		connection: db,
	}
}

func (db *orderConnection) Insert(order model.Order) model.Order {
	db.connection.Save(&order)
	return order
}

func (db *orderConnection) Update(order model.Order) model.Order {
	db.connection.Save(&order)
	return order
}

func (db *orderConnection) All() []model.Order {
	var orders []model.Order
	db.connection.Find(&orders)
	return orders
}

func (db *orderConnection) FindByID(orderID uint) model.Order {
	var order model.Order
	db.connection.Find(&order, orderID)
	return order
}

func (db *orderConnection) Delete(order model.Order) {
	db.connection.Delete(&order)
}
