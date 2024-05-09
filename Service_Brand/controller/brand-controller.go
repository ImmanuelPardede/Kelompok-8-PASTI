package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iqbalsiagian17/Service_Brand/dto"
	"github.com/iqbalsiagian17/Service_Brand/helper"
	"github.com/iqbalsiagian17/Service_Brand/model"
	"github.com/iqbalsiagian17/Service_Brand/service"
)

// brandController is a contract about something that this controller can do
type BrandController interface {
	All(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type brandController struct {
	BrandService service.BrandService
}

// NewbrandController creates a new instance of brandController
func NewBrandController(BrandService service.BrandService) BrandController {
	return &brandController{
		BrandService: BrandService,
	}
}

func (c *brandController) All(ctx *gin.Context) {
	brands := c.BrandService.All()
	ctx.JSON(http.StatusOK, brands)
}

func (c *brandController) FindByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	brand := c.BrandService.FindByID(id)
	if brand.ID == 0 {
		res := helper.BuildErrorResponse("Data not found", "No data with given ID", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, brand)
}

func (c *brandController) Insert(ctx *gin.Context) {
	var brandCreateDTO dto.BrandCreateDTO
	errDTO := ctx.ShouldBind(&brandCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result := c.BrandService.Insert(brandCreateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusCreated, response)
}

func (c *brandController) Update(ctx *gin.Context) {
	var brandUpdateDTO dto.BrandUpdateDTO
	errDTO := ctx.ShouldBind(&brandUpdateDTO)
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
	brandUpdateDTO.ID = uint(id) // Convert id to uint
	result := c.BrandService.Update(brandUpdateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *brandController) Delete(ctx *gin.Context) {
	var brand model.Brand
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	brand.ID = uint(id)
	c.BrandService.Delete(brand)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
