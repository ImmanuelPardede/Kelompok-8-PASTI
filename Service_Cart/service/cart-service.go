package service

import (
	"log"

	"github.com/iqbalsiagian17/Service_Cart/dto"
	"github.com/iqbalsiagian17/Service_Cart/model"
	"github.com/iqbalsiagian17/Service_Cart/repository"
	"github.com/mashingan/smapping"
)

// CartService is a contract about something that this service can do
type CartService interface {
	Insert(b dto.CartCreateDTO) model.Cart
	Update(b dto.CartUpdateDTO) model.Cart
	Delete(b model.Cart)
	All() []model.Cart
	FindByID(cartID uint64) model.Cart
}

type cartService struct {
	cartRepository repository.CartRepository
}

// NewCartService creates a new instance of CartService
func NewCartService(cartRepository repository.CartRepository) CartService {
	return &cartService{
		cartRepository: cartRepository,
	}
}

func (service *cartService) All() []model.Cart {
	return service.cartRepository.All()
}

func (service *cartService) FindByID(cartID uint64) model.Cart {
	id := uint(cartID)
	return service.cartRepository.FindByID(id)
}

func (service *cartService) Insert(b dto.CartCreateDTO) model.Cart {
	cart := model.Cart{}
	err := smapping.FillStruct(&cart, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.cartRepository.InsertCart(cart)
	return res
}

func (service *cartService) Update(b dto.CartUpdateDTO) model.Cart {
	cart := model.Cart{}
	err := smapping.FillStruct(&cart, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.cartRepository.UpdateCart(cart)
	return res
}

func (service *cartService) Delete(b model.Cart) {
	service.cartRepository.DeleteCart(b)
}
