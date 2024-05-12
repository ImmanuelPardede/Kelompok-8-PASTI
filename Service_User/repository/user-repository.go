package repository

import (
	"github.com/iqbalsiagian17/User/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) Create(user model.User) (model.User, error) {
	if err := ur.db.Create(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (ur *UserRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
