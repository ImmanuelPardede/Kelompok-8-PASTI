package repository

import (
	"github.com/iqbalsiagian17/carousel_service/model"
	"gorm.io/gorm"
)

// CarouselRepository adalah kontrak untuk mengelola Carousel.
type CarouselRepository interface {
	Create(item model.Carousel) model.Carousel
	Update(item model.Carousel) model.Carousel
	Index() []model.Carousel
	Show(itemID uint64) model.Carousel
	Delete(item model.Carousel)
}

type carouselConnection struct {
	connection *gorm.DB
}

// NewCarouselRepository membuat instance baru dari CarouselRepository.
func NewCarouselRepository(db *gorm.DB) CarouselRepository {
	return &carouselConnection{
		connection: db,
	}
}

func (db *carouselConnection) Create(item model.Carousel) model.Carousel {
	db.connection.Save(&item)
	return item
}

func (db *carouselConnection) Update(item model.Carousel) model.Carousel {
	db.connection.Save(&item)
	return item
}

func (db *carouselConnection) Index() []model.Carousel {
	var items []model.Carousel
	db.connection.Find(&items)
	return items
}

func (db *carouselConnection) Show(itemID uint64) model.Carousel {
	var item model.Carousel
	db.connection.Find(&item, itemID)
	return item
}

func (db *carouselConnection) Delete(item model.Carousel) {
	db.connection.Delete(&item)
}
