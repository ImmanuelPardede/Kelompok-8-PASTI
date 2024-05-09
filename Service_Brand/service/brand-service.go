package service

import (
	"log"

	"github.com/iqbalsiagian17/Service_Brand/dto"
	"github.com/iqbalsiagian17/Service_Brand/model"
	"github.com/iqbalsiagian17/Service_Brand/repository"
	"github.com/mashingan/smapping"
)

// BrandService is a contract about something that this service can do
type BrandService interface {
	Insert(b dto.BrandCreateDTO) model.Brand
	Update(b dto.BrandUpdateDTO) model.Brand
	Delete(b model.Brand)
	All() []model.Brand
	FindByID(brandID uint64) model.Brand
}

type brandService struct {
	brandRepository repository.BrandRepository
}

// NewBrandService creates a new instance of BrandService
func NewBrandService(brandRepository repository.BrandRepository) BrandService {
	return &brandService{
		brandRepository: brandRepository,
	}
}

func (service *brandService) All() []model.Brand {
	return service.brandRepository.All()
}

func (service *brandService) FindByID(brandID uint64) model.Brand {

	id := uint(brandID)
	return service.brandRepository.FindByID(id)
}

func (service *brandService) Insert(b dto.BrandCreateDTO) model.Brand {
	brand := model.Brand{}
	err := smapping.FillStruct(&brand, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.brandRepository.InsertBrand(brand)
	return res
}

func (service *brandService) Update(b dto.BrandUpdateDTO) model.Brand {
	brand := model.Brand{}
	err := smapping.FillStruct(&brand, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.brandRepository.UpdateBrand(brand)
	return res
}

func (service *brandService) Delete(b model.Brand) {
	service.brandRepository.DeleteBrand(b)
}
