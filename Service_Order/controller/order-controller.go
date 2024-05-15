package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iqbalsiagian17/Service_Order/dto"
	"github.com/iqbalsiagian17/Service_Order/helper"
	"github.com/iqbalsiagian17/Service_Order/service"
)

type OrderController interface {
	GetAllOrders(ctx *gin.Context)
	FindOrderByID(ctx *gin.Context)
	InsertOrder(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
	DeleteOrder(ctx *gin.Context)
}

type OrderControllerImpl struct {
	orderService   service.OrderService
	addressService service.AddressService
}

func NewOrderController(orderService service.OrderService, addressService service.AddressService) OrderController {
	return &OrderControllerImpl{
		orderService:   orderService,
		addressService: addressService,
	}
}

func (c *OrderControllerImpl) GetAllOrders(ctx *gin.Context) {
	orders := c.orderService.All()
	ctx.JSON(http.StatusOK, orders)
}

func (c *OrderControllerImpl) FindOrderByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	order := c.orderService.FindByID(id)
	if order.ID == 0 {
		res := helper.BuildErrorResponse("Data not found", "No data with given ID", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (c *OrderControllerImpl) InsertOrder(ctx *gin.Context) {
	var orderCreateDTO dto.OrderCreateDTO
	errDTO := ctx.ShouldBind(&orderCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Ambil UserID dari database menggunakan AddressID yang diberikan
	userID, err := c.addressService.GetUserIDByAddressID(orderCreateDTO.AddressID)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get UserID", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	// Create a new dto.OrderCreateDTO instance and populate it
	newOrderCreateDTO := dto.OrderCreateDTO{
		UserID:         userID,
		OrderDate:      orderCreateDTO.OrderDate,
		Total:          orderCreateDTO.Total,
		PaymentMethod:  orderCreateDTO.PaymentMethod,
		ShippingMethod: orderCreateDTO.ShippingMethod,
		AddressID:      orderCreateDTO.AddressID,
		Status:         orderCreateDTO.Status,
	}

	result := c.orderService.Insert(newOrderCreateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusCreated, response)
}

func (c *OrderControllerImpl) UpdateOrder(ctx *gin.Context) {
	var updateDTO dto.OrderUpdateDTO
	errDTO := ctx.ShouldBind(&updateDTO)
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
	updateDTO.ID = id

	// Check if order with given ID exists
	existingOrder := c.orderService.FindByID(id)
	if existingOrder.ID == 0 {
		res := helper.BuildErrorResponse("Data not found", "No data with given ID", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	updatedOrderDTO := dto.OrderUpdateDTO{
		ID:             updateDTO.ID,
		UserID:         updateDTO.UserID,
		OrderDate:      updateDTO.OrderDate,
		Total:          updateDTO.Total,
		PaymentMethod:  updateDTO.PaymentMethod,
		ShippingMethod: updateDTO.ShippingMethod,
		AddressID:      updateDTO.AddressID,
		Status:         updateDTO.Status,
	}

	result := c.orderService.Update(updatedOrderDTO)
	response := helper.BuildResponse(true, "Order updated successfully", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *OrderControllerImpl) DeleteOrder(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Check if order with given ID exists
	existingOrder := c.orderService.FindByID(id)
	if existingOrder.ID == 0 {
		res := helper.BuildErrorResponse("Data not found", "No data with given ID", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	c.orderService.Delete(existingOrder)
	res := helper.BuildResponse(true, "Order deleted successfully", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
