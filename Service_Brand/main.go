package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iqbalsiagian17/Service_Brand/config"
	"github.com/iqbalsiagian17/Service_Brand/controller"
	"github.com/iqbalsiagian17/Service_Brand/repository"
	"github.com/iqbalsiagian17/Service_Brand/service"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB                   = config.SetupDatabaseConnection()
	brandRepository repository.BrandRepository = repository.NewBrandRepository(db)
	BrandService    service.BrandService       = service.NewBrandService(brandRepository)
	brandController controller.BrandController = controller.NewBrandController(BrandService)
)

// membuat variable db dengan nilai setup database connection
func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	brandRoutes := r.Group("/api/brand")
	{
		brandRoutes.GET("/", brandController.All)
		brandRoutes.POST("/", brandController.Insert)
		brandRoutes.GET("/:id", brandController.FindByID)
		brandRoutes.PUT("/:id", brandController.Update)
		brandRoutes.DELETE("/:id", brandController.Delete)
	}
	r.Run(":9090")
}
