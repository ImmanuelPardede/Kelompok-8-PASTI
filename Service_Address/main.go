package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iqbalsiagian17/Service_Address/config"
	"github.com/iqbalsiagian17/Service_Address/controller"
	"github.com/iqbalsiagian17/Service_Address/repository"
	"github.com/iqbalsiagian17/Service_Address/service"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                     = config.SetupDatabaseConnection()
	addressRepository repository.AddressRepository = repository.NewAddressRepository(db)
	AddressService    service.AddressService       = service.NewAddressService(addressRepository)
	addressController controller.AddressController = controller.NewAddressController(AddressService)
)

// membuat variable db dengan nilai setup database connection
func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	addressRoutes := r.Group("/api/address")
	{
		addressRoutes.GET("/", addressController.All)
		addressRoutes.POST("/", addressController.Insert)
		addressRoutes.PUT("/:id", addressController.Update)
		addressRoutes.GET("/:id", addressController.FindByID)
		addressRoutes.DELETE("/:id", addressController.Delete)

	}
	r.Run(":9999")
}
