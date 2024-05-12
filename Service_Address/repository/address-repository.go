package repository

import (
	"github.com/iqbalsiagian17/Service_Address/model"
	"gorm.io/gorm"
)

type AddressRepository struct {
	DB *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{DB: db}
}

func (ar *AddressRepository) CreateAddress(address *model.Address) (*model.Address, error) {
	if err := ar.DB.Create(address).Error; err != nil {
		return nil, err
	}
	return address, nil
}

func (ar *AddressRepository) UpdateAddress(address *model.Address) (*model.Address, error) {
	if err := ar.DB.Save(address).Error; err != nil {
		return nil, err
	}
	return address, nil
}

func (ar *AddressRepository) DeleteAddress(address *model.Address) error {
	if err := ar.DB.Delete(address).Error; err != nil {
		return err
	}
	return nil
}

func (ar *AddressRepository) GetAddressByID(id uint) (*model.Address, error) {
	var address model.Address
	if err := ar.DB.First(&address, id).Error; err != nil {
		return nil, err
	}
	return &address, nil
}

func (ar *AddressRepository) GetAllAddresses() ([]*model.Address, error) {
	var addresses []*model.Address
	if err := ar.DB.Find(&addresses).Error; err != nil {
		return nil, err
	}
	return addresses, nil
}
