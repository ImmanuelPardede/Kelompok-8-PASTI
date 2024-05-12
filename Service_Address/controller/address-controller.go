package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iqbalsiagian17/Service_Address/dto"
	"github.com/iqbalsiagian17/Service_Address/helper"
	"github.com/iqbalsiagian17/Service_Address/model"
	"github.com/iqbalsiagian17/Service_Address/service"
)

// CategoryController is a contract about something that this controller can do
type AddressController interface {
	All(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

// CategoryService is a contract about something that this service can do
type UserService interface {
	GetUserID(id uint64) (uint64, error)
}

type userService struct{}

// NewCategoryService creates a new instance of CategoryService
func NewUserService() UserService {
	return &userService{}
}

type addressController struct {
	addressService service.AddressService
	userService    UserService
}

// NewAddressController creates a new instance of AddressController
func NewAddressController(AddressService service.AddressService, UserService UserService) AddressController {
	return &addressController{
		addressService: AddressService,
		userService:    UserService,
	}
}

func (c *addressController) All(ctx *gin.Context) {
	addresses := c.addressService.All()
	ctx.JSON(http.StatusOK, addresses)
}

func (c *addressController) FindByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	address := c.addressService.FindByID(id)
	if address.ID == 0 {
		res := helper.BuildErrorResponse("Data not found", "No data with given ID", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, address)
}

func (c *addressController) Insert(ctx *gin.Context) {
	var addressCreateDTO dto.AddressCreateDTO
	errDTO := ctx.ShouldBind(&addressCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Mendapatkan ID kategori dari layanan kategori

	userID, err := c.userService.GetUserID(uint64(addressCreateDTO.UserID))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get User ID", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	// Menambahkan ID kategori ke dalam addressCreateDTO
	addressCreateDTO.UserID = uint(userID)

	result := c.addressService.Insert(addressCreateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusCreated, response)
}

func (c *addressController) Update(ctx *gin.Context) {
	var addressUpdateDTO dto.AddressUpdateDTO
	errDTO := ctx.ShouldBind(&addressUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	idStr := ctx.Param("id")
	id, errID := strconv.ParseUint(idStr, 10, 64)
	if errID != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	addressUpdateDTO.ID = uint(id) // Convert id to uint

	// Mendapatkan ID kategori dari layanan kategori
	userID, err := c.userService.GetUserID(uint64(addressUpdateDTO.UserID))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get User ID", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	// Menambahkan ID kategori ke dalam AddressUpdateDTO
	addressUpdateDTO.UserID = uint(userID)

	result := c.addressService.Update(addressUpdateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *addressController) Delete(ctx *gin.Context) {
	var address model.Address
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	address.ID = uint(id)
	c.addressService.Delete(address)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (cs *userService) GetUserID(id uint64) (uint64, error) {
	// Panggil API kategori untuk mendapatkan informasi kategori berdasarkan ID
	url := fmt.Sprintf("http://localhost:8000/api/user/%d", id)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch User ID: %s", resp.Status)
	}

	var user struct {
		ID uint64 `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return 0, err
	}

	return user.ID, nil
}
