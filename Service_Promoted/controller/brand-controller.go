package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iqbalsiagian17/Service_Promoted/dto"
	"github.com/iqbalsiagian17/Service_Promoted/helper"
	"github.com/iqbalsiagian17/Service_Promoted/model"
	"github.com/iqbalsiagian17/Service_Promoted/service"
)

// PromotedController is a contract about something that this controller can do
type PromotedController interface {
	All(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type promotedController struct {
	PromotedService service.PromotedService
}

// NewPromotedController creates a new instance of PromotedController
func NewPromotedController(PromotedService service.PromotedService) PromotedController {
	return &promotedController{
		PromotedService: PromotedService,
	}
}

func (c *promotedController) All(ctx *gin.Context) {
	promoteds := c.PromotedService.All()
	ctx.JSON(http.StatusOK, promoteds)
}

func (c *promotedController) FindByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	promoted := c.PromotedService.FindByID(id)
	if promoted.ID == 0 {
		res := helper.BuildErrorResponse("Data not found", "No data with given ID", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, promoted)
}

func (c *promotedController) Insert(ctx *gin.Context) {
	var promotedCreateDTO dto.PromotedCreateDTO
	errDTO := ctx.ShouldBind(&promotedCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result := c.PromotedService.Insert(promotedCreateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusCreated, response)
}

func (c *promotedController) Update(ctx *gin.Context) {
	var promotedUpdateDTO dto.PromotedUpdateDTO
	errDTO := ctx.ShouldBind(&promotedUpdateDTO)
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
	promotedUpdateDTO.ID = uint(id) // Convert id to uint
	result := c.PromotedService.Update(promotedUpdateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *promotedController) Delete(ctx *gin.Context) {
	var promoted model.Promoted
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	promoted.ID = uint(id)
	c.PromotedService.Delete(promoted)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
