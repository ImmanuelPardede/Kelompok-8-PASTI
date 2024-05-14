package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iqbalsiagian17/Service_Product/config"
	"github.com/iqbalsiagian17/Service_Product/controller"
	"github.com/iqbalsiagian17/Service_Product/repository"
	"github.com/iqbalsiagian17/Service_Product/service"
	"gorm.io/gorm"
)

var (
	db                 *gorm.DB                     = config.SetupDatabaseConnection()
	productRepository  repository.ProductRepository = repository.NewProductRepository(db)
	categoryService    service.CategoryService      = service.NewCategoryService()
	subcategoryService service.SubCategoryService   = service.NewSubCategoryService()
	brandService       service.BrandService         = service.NewBrandService()
	ProductService     service.ProductService       = service.NewProductService(productRepository, categoryService, subcategoryService, brandService)
	productController  controller.ProductController = controller.NewProductController(ProductService, categoryService, subcategoryService, brandService)
)

// membuat variable db dengan nilai setup database connection
func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	productRoutes := r.Group("/api/product")
	{
		productRoutes.GET("/", productController.All)
		productRoutes.POST("/", productController.Insert)
		productRoutes.GET("/:id", productController.FindByID)
		productRoutes.PUT("/:id", productController.Update)
		productRoutes.DELETE("/:id", productController.Delete)
	}
	r.Run(":2222")
}
