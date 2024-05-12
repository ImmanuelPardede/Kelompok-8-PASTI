package service

import (
	"encoding/json" // Ini adalah impor yang diperlukan
	"fmt"
	"log"
	"net/http"

	"github.com/iqbalsiagian17/Service_Product/dto"
	"github.com/iqbalsiagian17/Service_Product/model"
	"github.com/iqbalsiagian17/Service_Product/repository"
	"github.com/mashingan/smapping"
)

// productService is a contract about something that this service can do
type ProductService interface {
	Insert(b dto.ProductCreateDTO) model.Product
	Update(b dto.ProductUpdateDTO) model.Product
	Delete(b model.Product)
	All() []model.Product
	FindByID(addressID uint64) model.Product
}

type productService struct {
	productRepository repository.ProductRepository
}

type CategoryService interface {
	GetCategoryID(id uint64) (uint64, error)
}

type categoryService struct{}

func NewCategoryService() CategoryService {
	return &categoryService{}
}

type SubCategoryService interface {
	GetSubCategoryID(id uint64) (uint64, error)
}

type subcategoryService struct{}

func NewSubCategoryService() SubCategoryService {
	return &subcategoryService{}
}

type BrandService interface {
	GetBrandID(id uint64) (uint64, error)
}

type brandService struct{}

func NewBrandService() BrandService {
	return &brandService{}
}

// NewProductService creates a new instance of ProductService
func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepository,
	}
}

func (service *productService) All() []model.Product {
	return service.productRepository.All()
}

func (service *productService) FindByID(productID uint64) model.Product {

	id := uint(productID)
	return service.productRepository.FindByID(id)
}

func (service *productService) Insert(b dto.ProductCreateDTO) model.Product {
	product := model.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	// Set category_id from DTO to the SubCategory model
	product.CategoryID = b.CategoryID
	product.SubCategoryID = b.SubCategoryID
	product.BrandID = b.BrandID

	res := service.productRepository.InsertProduct(product)
	return res
}

func (service *productService) Update(b dto.ProductUpdateDTO) model.Product {
	product := model.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	// Set category_id from DTO to the SubCategory model
	product.CategoryID = b.CategoryID
	product.SubCategoryID = b.SubCategoryID
	product.BrandID = b.BrandID

	res := service.productRepository.UpdateProduct(product)
	return res
}

func (service *productService) Delete(b model.Product) {
	service.productRepository.DeleteProduct(b)
}

func (cs *categoryService) GetCategoryID(id uint64) (uint64, error) {
	// Replace the URL with the actual endpoint of your category service API
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
	// Replace the URL with the actual endpoint of your category service API
	url := fmt.Sprintf("http://localhost:8888/api/subcategory/%d", id)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch SubCategory ID: %s", resp.Status)
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
	// Replace the URL with the actual endpoint of your category service API
	url := fmt.Sprintf("http://localhost:6666/api/brand/%d", id)

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
