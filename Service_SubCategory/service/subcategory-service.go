package service

import (
	"encoding/json" // Ini adalah impor yang diperlukan
	"fmt"
	"log"
	"net/http"

	"github.com/iqbalsiagian17/Service_SubCategory/dto"
	"github.com/iqbalsiagian17/Service_SubCategory/model"
	"github.com/iqbalsiagian17/Service_SubCategory/repository"
	"github.com/mashingan/smapping"
)

// SubcategoryService is a contract about something that this service can do
type SubcategoryService interface {
	Insert(b dto.SubCategoryCreateDTO) model.SubCategory
	Update(b dto.SubCategoryUpdateDTO) model.SubCategory
	Delete(b model.SubCategory)
	All() []model.SubCategory
	FindByID(categoryID uint64) model.SubCategory
}

type subcategoryService struct {
	subcategoryRepository repository.SubcategoryRepository
}

type CategoryService interface {
	GetCategoryID(id uint64) (uint64, error)
}

type categoryService struct{}

func NewCategoryService() CategoryService {
	return &categoryService{}
}

// NewSubcategoryService creates a new instance of SubcategoryService
func NewSubcategoryService(subcategoryRepository repository.SubcategoryRepository) SubcategoryService {
	return &subcategoryService{
		subcategoryRepository: subcategoryRepository,
	}
}

func (service *subcategoryService) All() []model.SubCategory {
	return service.subcategoryRepository.All()
}

func (service *subcategoryService) FindByID(subcategoryID uint64) model.SubCategory {

	id := uint(subcategoryID)
	return service.subcategoryRepository.FindByID(id)
}

func (service *subcategoryService) Insert(b dto.SubCategoryCreateDTO) model.SubCategory {
	subcategory := model.SubCategory{}
	err := smapping.FillStruct(&subcategory, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	// Set category_id from DTO to the SubCategory model
	subcategory.CategoryID = b.CategoryID

	res := service.subcategoryRepository.InsertSubcategory(subcategory)
	return res
}

func (service *subcategoryService) Update(b dto.SubCategoryUpdateDTO) model.SubCategory {
	subcategory := model.SubCategory{}
	err := smapping.FillStruct(&subcategory, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	// Set category_id from DTO to the SubCategory model
	subcategory.CategoryID = b.CategoryID

	res := service.subcategoryRepository.UpdateSubcategory(subcategory)
	return res
}

func (service *subcategoryService) Delete(b model.SubCategory) {
	service.subcategoryRepository.DeleteSubcategory(b)
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
