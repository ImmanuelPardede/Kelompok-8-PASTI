package repository

import (
	"github.com/iqbalsiagian17/Service_Promoted/model"
	"gorm.io/gorm"
)

type PromotedRepository interface {
	InsertPromoted(promoted model.Promoted) model.Promoted
	UpdatePromoted(promoted model.Promoted) model.Promoted
	All() []model.Promoted
	FindByID(PromotedID uint) model.Promoted
	DeletePromoted(promoted model.Promoted)
}

type promotedConnection struct {
	connection *gorm.DB
}

func NewPromotedRepository(db *gorm.DB) PromotedRepository {
	return &promotedConnection{
		connection: db,
	}
}

func (db *promotedConnection) InsertPromoted(promoted model.Promoted) model.Promoted {
	db.connection.Save(&promoted)
	return promoted
}

func (db *promotedConnection) UpdatePromoted(promoted model.Promoted) model.Promoted {
	db.connection.Save(&promoted)
	return promoted
}

func (db *promotedConnection) All() []model.Promoted {
	var promoteds []model.Promoted
	db.connection.Find(&promoteds)
	return promoteds
}

func (db *promotedConnection) FindByID(promotedID uint) model.Promoted {
	var promoted model.Promoted
	db.connection.Find(&promoted, promotedID)
	return promoted
}

func (db *promotedConnection) DeletePromoted(promoted model.Promoted) {
	db.connection.Delete(&promoted)
}
