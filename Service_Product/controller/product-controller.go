package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iqbalsiagian17/Service_Product/dto"
	"github.com/iqbalsiagian17/Service_Product/helper"
	"github.com/iqbalsiagian17/Service_Product/model"
	"github.com/iqbalsiagian17/Service_Product/service"
)

// CategoryController is a contract about something that this controller can do
type ProductController interface {
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

type SubCategoryService interface {
	GetSubCategoryID(id uint64) (uint64, error)
}

type subcategoryService struct{}

// NewSCategoryService creates a new instance of CategoryService
func NewSubCategoryService() SubCategoryService {
	return &subcategoryService{}
}

type BrandService interface {
	GetBrandID(id uint64) (uint64, error)
}

type brandService struct{}

// NewSCategoryService creates a new instance of CategoryService
func NewBrandService() BrandService {
	return &brandService{}
}

type productController struct {
	productService     service.ProductService
	categoryService    CategoryService
	subcategoryService SubCategoryService
	brandService       BrandService
}

// NewSubCategoryController creates a new instance of SubcategoryController
func NewProductController(ProductService service.ProductService, CategoryService CategoryService, SubcategoryService SubCategoryService, BrandService BrandService) ProductController {
	return &productController{
		productService:     ProductService,
		categoryService:    CategoryService,
		subcategoryService: SubcategoryService,
		brandService:       BrandService,
	}
}

func (c *productController) All(ctx *gin.Context) {
	products := c.productService.All()
	ctx.JSON(http.StatusOK, products)
}

func (c *productController) FindByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	product := c.productService.FindByID(id)
	if product.ID == 0 {
		res := helper.BuildErrorResponse("Data not found", "No data with given ID", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (c *productController) Insert(ctx *gin.Context) {
	var productCreateDTO dto.ProductCreateDTO
	errDTO := ctx.ShouldBind(&productCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Mendapatkan ID kategori dari layanan kategori
	categoryID, err := c.categoryService.GetCategoryID(uint64(productCreateDTO.CategoryID))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get category ID", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	// Menambahkan ID kategori ke dalam productCreateDTO
	productCreateDTO.CategoryID = uint(categoryID)

	// Mendapatkan ID subkategori dari layanan subkategori
	subcategoryID, err := c.subcategoryService.GetSubCategoryID(uint64(productCreateDTO.SubCategoryID))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get Subcategory ID", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	// Menambahkan ID subkategori ke dalam productCreateDTO
	productCreateDTO.SubCategoryID = uint(subcategoryID)

	// Mendapatkan ID Brand dari layanan brand
	brandID, err := c.brandService.GetBrandID(uint64(productCreateDTO.BrandID))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get Brand ID", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	// Menambahkan ID Brand ke dalam productCreateDTO
	productCreateDTO.BrandID = uint(brandID)

	result := c.productService.Insert(productCreateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusCreated, response)
}

func (c *productController) Update(ctx *gin.Context) {
	var productUpdateDTO dto.ProductUpdateDTO
	errDTO := ctx.ShouldBind(&productUpdateDTO)
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
	productUpdateDTO.ID = uint(id) // Convert id to uint

	// Mendapatkan ID kategori dari layanan kategori
	categoryID, err := c.categoryService.GetCategoryID(uint64(productUpdateDTO.CategoryID))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get category ID", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	// Menambahkan ID kategori ke dalam productUpdateDTO
	productUpdateDTO.CategoryID = uint(categoryID)

	// Mendapatkan ID subkategori dari layanan kategori
	subcategoryID, err := c.subcategoryService.GetSubCategoryID(uint64(productUpdateDTO.SubCategoryID))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get Subcategory ID", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	// Menambahkan ID subkategori ke dalam productUpdateDTO
	productUpdateDTO.SubCategoryID = uint(subcategoryID)

	// Mendapatkan ID brand dari layanan kategori
	brandID, err := c.brandService.GetBrandID(uint64(productUpdateDTO.BrandID))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get Brand ID", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	// Menambahkan ID brand ke dalam productUpdateDTO
	productUpdateDTO.BrandID = uint(brandID)

	result := c.productService.Update(productUpdateDTO)
	response := helper.BuildResponse(true, "OK!", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *productController) Delete(ctx *gin.Context) {
	var product model.Product
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get ID", "No param ID were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	product.ID = uint(id)
	c.productService.Delete(product)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (cs *categoryService) GetCategoryID(id uint64) (uint64, error) {
	url := fmt.Sprintf("http://localhost:7777/api/category/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch Category ID: %s", resp.Status)
	}

	var category struct {
		ID uint64 `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&category); err != nil {
		return 0, err
	}

	return category.ID, nil
}

func (cs *subcategoryService) GetSubCategoryID(id uint64) (uint64, error) {
	// Panggil API kategori untuk mendapatkan informasi kategori berdasarkan ID
	url := fmt.Sprintf("http://localhost:8888/api/subcategory/%d", id)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch Category ID: %s", resp.Status)
	}

	var subcategory struct {
		ID uint64 `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&subcategory); err != nil {
		return 0, err
	}

	return subcategory.ID, nil
}

func (cs *brandService) GetBrandID(id uint64) (uint64, error) {
	// Panggil API kategori untuk mendapatkan informasi kategori berdasarkan ID
	url := fmt.Sprintf("http://localhost:9090/api/brand/%d", id)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch Brand ID: %s", resp.Status)
	}

	var brand struct {
		ID uint64 `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&brand); err != nil {
		return 0, err
	}

	return brand.ID, nil
}
