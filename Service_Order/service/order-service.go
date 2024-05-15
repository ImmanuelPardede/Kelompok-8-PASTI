package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/iqbalsiagian17/Service_Order/dto"
	"github.com/iqbalsiagian17/Service_Order/model"
	"github.com/iqbalsiagian17/Service_Order/repository"
	"github.com/mashingan/smapping"
)

type AddressService interface {
	GetAddressID(addressID uint64) (uint64, error)
	GetUserIDByAddressID(addressID uint64) (uint64, error)
}

type OrderService interface {
	Insert(b dto.OrderCreateDTO) model.Order
	Update(b dto.OrderUpdateDTO) model.Order
	Delete(b model.Order)
	All() []model.Order
	FindByID(orderID uint64) model.Order
}

type AddressServiceImpl struct{}

func (as *AddressServiceImpl) GetAddressID(id uint64) (uint64, error) {
	url := fmt.Sprintf("http://localhost:9999/api/address/%d", id)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch Address ID: %s", resp.Status)
	}

	var address struct {
		ID uint64 `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
		return 0, err
	}

	return address.ID, nil
}

func (as *AddressServiceImpl) GetUserIDByAddressID(addressID uint64) (uint64, error) {
	url := fmt.Sprintf("http://localhost:9999/api/address/%d/userid/%d", addressID, addressID)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch User ID: %s", resp.Status)
	}

	var userID uint64
	if err := json.NewDecoder(resp.Body).Decode(&userID); err != nil {
		return 0, err
	}

	return userID, nil
}

type OrderServiceImpl struct {
	orderRepository repository.OrderRepository
	addressService  AddressService
}

func NewAddressService() AddressService {
	return &AddressServiceImpl{}
}

func NewOrderService(orderRepository repository.OrderRepository, addressService AddressService) OrderService {
	return &OrderServiceImpl{
		orderRepository: orderRepository,
		addressService:  addressService,
	}
}

func (service *OrderServiceImpl) All() []model.Order {
	return service.orderRepository.All()
}

func (service *OrderServiceImpl) FindByID(orderID uint64) model.Order {
	id := uint(orderID)
	return service.orderRepository.FindByID(id)
}

func (service *OrderServiceImpl) Insert(b dto.OrderCreateDTO) model.Order {
	order := model.Order{}
	err := smapping.FillStruct(&order, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.orderRepository.Insert(order)
	return res
}

func (service *OrderServiceImpl) Update(b dto.OrderUpdateDTO) model.Order {
	order := model.Order{}
	err := smapping.FillStruct(&order, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	// Set category_id from DTO to the Address model
	order.UserID = b.UserID

	res := service.orderRepository.Update(order)
	return res
}

func (service *OrderServiceImpl) Delete(b model.Order) {
	service.orderRepository.Delete(b)
}
