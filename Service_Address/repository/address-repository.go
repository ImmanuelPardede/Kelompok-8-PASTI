package repository

import (
	"github.com/iqbalsiagian17/Service_Address/model"
	"gorm.io/gorm"
)

type AddressRepository interface {
	InsertAddress(address model.Address) model.Address
	UpdateAddress(address model.Address) model.Address
	All() []model.Address
	FindByID(AddressID uint) model.Address
	DeleteAddress(address model.Address)
	FindByUserID(userID uint) []model.Address // Add this line
}

type addressConnection struct {
	connection *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressConnection{
		connection: db,
	}
}

func (db *addressConnection) InsertAddress(address model.Address) model.Address {
	db.connection.Save(&address)
	return address
}

func (db *addressConnection) UpdateAddress(address model.Address) model.Address {
	db.connection.Save(&address)
	return address
}

func (db *addressConnection) All() []model.Address {
	var addresses []model.Address
	db.connection.Find(&addresses)
	return addresses
}

func (db *addressConnection) FindByID(addressID uint) model.Address {
	var address model.Address
	db.connection.Find(&address, addressID)
	return address
}

func (db *addressConnection) FindByUserID(userID uint) []model.Address {
	var addresses []model.Address
	db.connection.Where("user_id = ?", userID).Find(&addresses)
	return addresses
}

func (db *addressConnection) DeleteAddress(address model.Address) {
	db.connection.Delete(&address)
}
