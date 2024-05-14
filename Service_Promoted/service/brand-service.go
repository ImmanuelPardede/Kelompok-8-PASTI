package service

import (
	"log"

	"github.com/iqbalsiagian17/Service_Promoted/dto"
	"github.com/iqbalsiagian17/Service_Promoted/model"
	"github.com/iqbalsiagian17/Service_Promoted/repository"
	"github.com/mashingan/smapping"
)

// PromotedService is a contract about something that this service can do
type PromotedService interface {
	Insert(b dto.PromotedCreateDTO) model.Promoted
	Update(b dto.PromotedUpdateDTO) model.Promoted
	Delete(b model.Promoted)
	All() []model.Promoted
	FindByID(promotedID uint64) model.Promoted
}

type promotedService struct {
	promotedRepository repository.PromotedRepository
}

// NewPromotedService creates a new instance of PromotedService
func NewPromotedService(promotedRepository repository.PromotedRepository) PromotedService {
	return &promotedService{
		promotedRepository: promotedRepository,
	}
}

func (service *promotedService) All() []model.Promoted {
	return service.promotedRepository.All()
}

func (service *promotedService) FindByID(promotedID uint64) model.Promoted {

	id := uint(promotedID)
	return service.promotedRepository.FindByID(id)
}

func (service *promotedService) Insert(b dto.PromotedCreateDTO) model.Promoted {
	promoted := model.Promoted{}
	err := smapping.FillStruct(&promoted, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.promotedRepository.InsertPromoted(promoted)
	return res
}

func (service *promotedService) Update(b dto.PromotedUpdateDTO) model.Promoted {
	promoted := model.Promoted{}
	err := smapping.FillStruct(&promoted, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.promotedRepository.UpdatePromoted(promoted)
	return res
}

func (service *promotedService) Delete(b model.Promoted) {
	service.promotedRepository.DeletePromoted(b)
}
