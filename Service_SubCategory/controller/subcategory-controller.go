package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iqbalsiagian17/Service_SubCategory/dto"
	"github.com/iqbalsiagian17/Service_SubCategory/helper"
	"github.com/iqbalsiagian17/Service_SubCategory/model"
	"github.com/iqbalsiagian17/Service_SubCategory/service"
)

// CategoryController is a contract about something that this controller can do
type SubcategoryController interface {
	All(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

// CategoryService is a contract about something that this service can do
type CategoryService interface {
	GetCategoryID(id uint64) (uint64, error)
}

type categoryService struct{}

// NewCategoryService creates a new instance of CategoryService
func NewCategoryService() CategoryService {
	return &categoryService{}
}

type subcategoryController struct {
	subCategoryService service.SubcategoryService
	categoryService    CategoryService
}

// NewSubCategoryController creates a new instance of SubcategoryController
func NewSubCategoryController(SubcategoryService service.SubcategoryService, CategoryService CategoryService) SubcategoryController {
	return &subcategoryController{
		subCategoryService: SubcategoryService,
		categoryService:    CategoryService,
	}
}

func (c *subcategoryController) All(ctx *gin.Context) {
	subcategories := c.subCategoryService.All()
	ctx.JSON(http.StatusOK, subcategories)
}

func (c *subcategoryController) FindByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	subcategory := c.subCategoryService.FindByID(id)
	if subcategory.ID == 0 {
		res := helper.BuildErrorResponse("Data not found", "No data with given ID", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, subcategory)
}

func (c *subcategoryController) Insert(ctx *gin.Context) {
	var subcategoryCreateDTO dto.SubCategoryCreateDTO
	errDTO := ctx.ShouldBind(&subcategoryCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Mendapatkan ID kategori dari layanan kategori
	categoryID, err := c.categoryService.GetCategoryID(uint64(subcategoryCreateDTO.CategoryID))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get category ID", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	// Menambahkan ID kategori ke dalam SubCategoryCreateDTO
	subcategoryCreateDTO.CategoryID = uint(categoryID)

	result := c.subCategoryService.Insert(subcategoryCreateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusCreated, response)
}

func (c *subcategoryController) Update(ctx *gin.Context) {
	var subcategoryUpdateDTO dto.SubCategoryUpdateDTO
	errDTO := ctx.ShouldBind(&subcategoryUpdateDTO)
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
	subcategoryUpdateDTO.ID = uint(id) // Convert id to uint

	// Mendapatkan ID kategori dari layanan kategori
	categoryID, err := c.categoryService.GetCategoryID(uint64(subcategoryUpdateDTO.CategoryID))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get category ID", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	// Menambahkan ID kategori ke dalam SubCategoryUpdateDTO
	subcategoryUpdateDTO.CategoryID = uint(categoryID)

	result := c.subCategoryService.Update(subcategoryUpdateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *subcategoryController) Delete(ctx *gin.Context) {
	var subcategory model.SubCategory
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	subcategory.ID = uint(id)
	c.subCategoryService.Delete(subcategory)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (cs *categoryService) GetCategoryID(id uint64) (uint64, error) {
	// Panggil API kategori untuk mendapatkan informasi kategori berdasarkan ID
	url := fmt.Sprintf("http://localhost:7777/api/category/%d", id)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch Category ID: %s", resp.Status)
	}

	var kategori struct {
		ID uint64 `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&kategori); err != nil {
		return 0, err
	}

	return kategori.ID, nil
}
