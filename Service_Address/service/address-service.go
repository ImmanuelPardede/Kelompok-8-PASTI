package service

import (
	"encoding/json" // Ini adalah impor yang diperlukan
	"fmt"
	"log"
	"net/http"

	"github.com/iqbalsiagian17/Service_Address/dto"
	"github.com/iqbalsiagian17/Service_Address/model"
	"github.com/iqbalsiagian17/Service_Address/repository"
	"github.com/mashingan/smapping"
)

// AddressService is a contract about something that this service can do
type AddressService interface {
	Insert(b dto.AddressCreateDTO) model.Address
	Update(b dto.AddressUpdateDTO) model.Address
	Delete(b model.Address)
	All() []model.Address
	FindByID(addressID uint64) model.Address
}

type addressService struct {
	addressRepository repository.AddressRepository
}

type UserService interface {
	GetUserID(id uint64) (uint64, error)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

// NewAddressService creates a new instance of AddressService
func NewAddressService(addressRepository repository.AddressRepository) addressService {
	return addressService{
		addressRepository: addressRepository,
	}
}

func (service addressService) All() []model.Address {
	return service.addressRepository.All()
}

func (service addressService) FindByID(addressID uint64) model.Address {

	id := uint(addressID)
	return service.addressRepository.FindByID(id)
}

func (service addressService) Insert(b dto.AddressCreateDTO) model.Address {
	address := model.Address{}
	err := smapping.FillStruct(&address, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	// Set user_id from DTO to the Address model
	address.UserID = b.UserID

	res := service.addressRepository.InsertAddress(address)
	return res
}

func (service addressService) Update(b dto.AddressUpdateDTO) model.Address {
	address := model.Address{}
	err := smapping.FillStruct(&address, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	// Set category_id from DTO to the Address model
	address.UserID = b.UserID

	res := service.addressRepository.UpdateAddress(address)
	return res
}

func (service addressService) Delete(b model.Address) {
	service.addressRepository.DeleteAddress(b)
}

func (cs *userService) GetUserID(id uint64) (uint64, error) {
	// Replace the URL with the actual endpoint of your category service API
	url := fmt.Sprintf("http://localhost:8000/api/user/%d", id)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch User id: %s", resp.Status)
	}

	var user struct {
		ID uint64 `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return 0, err
	}

	return user.ID, nil
}
