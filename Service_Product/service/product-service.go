package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/iqbalsiagian17/Service_Product/dto"
	"github.com/iqbalsiagian17/Service_Product/model"
	"github.com/iqbalsiagian17/Service_Product/repository"
	"github.com/mashingan/smapping"
)

type ProductService interface {
	Insert(b dto.ProductCreateDTO) model.Product
	Update(b dto.ProductUpdateDTO) model.Product
	Delete(b model.Product)
	All() []model.Product
	FindByID(productID uint64) model.Product
}

type productService struct {
	productRepository  repository.ProductRepository
	categoryService    CategoryService
	subcategoryService SubCategoryService
	brandService       BrandService
}

func NewProductService(productRepository repository.ProductRepository, categoryService CategoryService, subcategoryService SubCategoryService, brandService BrandService) ProductService {
	return &productService{
		productRepository:  productRepository,
		categoryService:    categoryService,
		subcategoryService: subcategoryService,
		brandService:       brandService,
	}
}

func (s *productService) All() []model.Product {
	return s.productRepository.All()
}

func (s *productService) FindByID(productID uint64) model.Product {
	// Konversi tipe data uint64 ke uint sebelum memanggil fungsi FindByID
	return s.productRepository.FindByID(uint(productID))
}

func (s *productService) Insert(b dto.ProductCreateDTO) model.Product {
	product := model.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	// Set category_id from DTO to the Product model
	categoryID, err := s.categoryService.GetCategoryID(uint64(b.CategoryID)) // Konversi ke uint64
	if err != nil {
		log.Fatalf("Failed to get category ID: %v", err)
	}
	product.CategoryID = uint(categoryID) // Konversi ke uint

	// Set subcategory_id from DTO to the Product model
	subcategoryID, err := s.subcategoryService.GetSubCategoryID(uint64(b.SubCategoryID)) // Konversi ke uint64
	if err != nil {
		log.Fatalf("Failed to get subcategory ID: %v", err)
	}
	product.SubCategoryID = uint(subcategoryID) // Konversi ke uint

	// Set brand_id from DTO to the Product model
	brandID, err := s.brandService.GetBrandID(uint64(b.BrandID)) // Konversi ke uint64
	if err != nil {
		log.Fatalf("Failed to get brand ID: %v", err)
	}
	product.BrandID = uint(brandID) // Konversi ke uint

	res := s.productRepository.InsertProduct(product)
	return res
}

func (s *productService) Update(b dto.ProductUpdateDTO) model.Product {
	product := model.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	// Set category_id from DTO to the Product model
	categoryID, err := s.categoryService.GetCategoryID(uint64(b.CategoryID)) // Konversi ke uint64
	if err != nil {
		log.Fatalf("Failed to get category ID: %v", err)
	}
	product.CategoryID = uint(categoryID) // Konversi ke uint

	// Set subcategory_id from DTO to the Product model
	subcategoryID, err := s.subcategoryService.GetSubCategoryID(uint64(b.SubCategoryID)) // Konversi ke uint64
	if err != nil {
		log.Fatalf("Failed to get subcategory ID: %v", err)
	}
	product.SubCategoryID = uint(subcategoryID) // Konversi ke uint

	// Set brand_id from DTO to the Product model
	brandID, err := s.brandService.GetBrandID(uint64(b.BrandID)) // Konversi ke uint64
	if err != nil {
		log.Fatalf("Failed to get brand ID: %v", err)
	}
	product.BrandID = uint(brandID) // Konversi ke uint

	res := s.productRepository.UpdateProduct(product)
	return res
}

func (s *productService) Delete(b model.Product) {
	s.productRepository.DeleteProduct(b)
}

type CategoryService interface {
	GetCategoryID(id uint64) (uint64, error)
}

type SubCategoryService interface {
	GetSubCategoryID(id uint64) (uint64, error)
}

type BrandService interface {
	GetBrandID(id uint64) (uint64, error)
}

// Implement CategoryService, SubCategoryService, and BrandService

type categoryService struct{}

func NewCategoryService() CategoryService {
	return &categoryService{}
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

// Implement SubCategoryService

type subcategoryService struct{}

func NewSubCategoryService() SubCategoryService {
	return &subcategoryService{}
}

func (cs *subcategoryService) GetSubCategoryID(id uint64) (uint64, error) {
	// Replace the URL with the actual endpoint of your subcategory service API
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

// Implement BrandService

type brandService struct{}

func NewBrandService() BrandService {
	return &brandService{}
}

func (cs *brandService) GetBrandID(id uint64) (uint64, error) {
	// Replace the URL with the actual endpoint of your brand service API
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
