package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iqbalsiagian17/carousel_service/dto"
	"github.com/iqbalsiagian17/carousel_service/helper"
	"github.com/iqbalsiagian17/carousel_service/model"
	"github.com/iqbalsiagian17/carousel_service/service"
)

// CarouselController adalah interface untuk mengelola permintaan terkait CarouselItem
type CarouselController interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

// CarouselController adalah controller untuk mengelola Carousel.
type carouselController struct {
	CarouselService service.CarouselService
}

// NewCarouselController creates a new instance of CarouselController
func NewCarouselController(carouselService service.CarouselService) CarouselController {
	return &carouselController{
		CarouselService: carouselService,
	}
}

func (ctrl *carouselController) Index(ctx *gin.Context) {
	categories := ctrl.CarouselService.Index()
	ctx.JSON(http.StatusOK, gin.H{"categories": categories})
}

func (c *carouselController) Show(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	carousel := c.CarouselService.Show(id)
	if (carousel == model.Carousel{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given ID", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK!", carousel)
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *carouselController) Create(ctx *gin.Context) {
	var carouselCreateDTO dto.CarouselCreateDTO
	errDTO := ctx.ShouldBind(&carouselCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result := c.CarouselService.Create(carouselCreateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusCreated, response)
}

func (c *carouselController) Update(ctx *gin.Context) {
	var carouselUpdateDTO dto.CarouselUpdateDTO
	errDTO := ctx.ShouldBind(&carouselUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	id, errID := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if errID != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	carouselUpdateDTO.ID = id
	result := c.CarouselService.Update(carouselUpdateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *carouselController) Delete(ctx *gin.Context) {
	var carousel model.Carousel
	id, errID := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if errID != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	carousel.ID = id
	c.CarouselService.Delete(carousel)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
