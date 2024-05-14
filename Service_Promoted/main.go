package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iqbalsiagian17/Service_Promoted/config"
	"github.com/iqbalsiagian17/Service_Promoted/controller"
	"github.com/iqbalsiagian17/Service_Promoted/repository"
	"github.com/iqbalsiagian17/Service_Promoted/service"
	"gorm.io/gorm"
)

var (
	db                 *gorm.DB                      = config.SetupDatabaseConnection()
	promotedRepository repository.PromotedRepository = repository.NewPromotedRepository(db)
	PromotedService    service.PromotedService       = service.NewPromotedService(promotedRepository)
	promotedController controller.PromotedController = controller.NewPromotedController(PromotedService)
)

// membuat variable db dengan nilai setup database connection
func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	promotedRoutes := r.Group("/api/promoted")
	{
		promotedRoutes.GET("/", promotedController.All)
		promotedRoutes.POST("/", promotedController.Insert)
		promotedRoutes.GET("/:id", promotedController.FindByID)
		promotedRoutes.PUT("/:id", promotedController.Update)
		promotedRoutes.DELETE("/:id", promotedController.Delete)
	}
	r.Run(":2020")
}
