package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iqbalsiagian17/Service_SubCategory/config"
	"github.com/iqbalsiagian17/Service_SubCategory/controller"
	"github.com/iqbalsiagian17/Service_SubCategory/repository"
	"github.com/iqbalsiagian17/Service_SubCategory/service"
	"gorm.io/gorm"
)

var (
	db                    *gorm.DB                         = config.SetupDatabaseConnection()
	subcategoryRepository repository.SubcategoryRepository = repository.NewSubCategoryRepository(db)
	SubcategoryService    service.SubcategoryService       = service.NewSubcategoryService(subcategoryRepository)
	categoryService service.CategoryService = service.NewCategoryService()
	subcategoryController controller.SubcategoryController = controller.NewSubCategoryController(SubcategoryService, categoryService)
)

// membuat variable db dengan nilai setup database connection
func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	subcategoryRoutes := r.Group("/api/subcategory")
	{
		subcategoryRoutes.GET("/", subcategoryController.All)
		subcategoryRoutes.POST("/", subcategoryController.Insert)
		subcategoryRoutes.GET("/:id", subcategoryController.FindByID)
		subcategoryRoutes.PUT("/:id", subcategoryController.Update)
		subcategoryRoutes.DELETE("/:id", subcategoryController.Delete)
	}
	r.Run(":8888")
}
