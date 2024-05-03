package service

import (
	"log"

	"github.com/iqbalsiagian17/carousel_service/dto"
	"github.com/iqbalsiagian17/carousel_service/model"
	"github.com/iqbalsiagian17/carousel_service/repository"
	"github.com/mashingan/smapping"
)

// CarouselService adalah kontrak tentang apa yang dapat dilakukan oleh layanan ini
type CarouselService interface {
	Create(b dto.CarouselCreateDTO) model.Carousel
	Update(b dto.CarouselUpdateDTO) model.Carousel
	Delete(b model.Carousel)
	Index() []model.Carousel
	Show(carouselID uint64) model.Carousel
}

type carouselService struct {
	carouselRepository repository.CarouselRepository
}

// NewCarouselService membuat instance baru dari CarouselService
func NewCarouselService(carouselRepository repository.CarouselRepository) CarouselService {
	return &carouselService{
		carouselRepository: carouselRepository,
	}
}

func (service *carouselService) Index() []model.Carousel {
	return service.carouselRepository.Index()
}

func (service *carouselService) Show(carouselID uint64) model.Carousel {
	return service.carouselRepository.Show(carouselID)
}

func (service *carouselService) Create(b dto.CarouselCreateDTO) model.Carousel {
	carousel := model.Carousel{}
	err := smapping.FillStruct(&carousel, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.carouselRepository.Create(carousel)
	return res
}

func (service *carouselService) Update(b dto.CarouselUpdateDTO) model.Carousel {
	product := model.Carousel{}
	err := smapping.FillStruct(&product, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.carouselRepository.Update(product)
	return res
}

func (service *carouselService) Delete(b model.Carousel) {
	service.carouselRepository.Delete(b)
}
