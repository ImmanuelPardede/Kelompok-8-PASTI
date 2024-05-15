package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iqbalsiagian17/Service_Order/config"
	"github.com/iqbalsiagian17/Service_Order/controller"
	"github.com/iqbalsiagian17/Service_Order/repository"
	"github.com/iqbalsiagian17/Service_Order/service"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB                   = config.SetupDatabaseConnection()
	orderRepository repository.OrderRepository = repository.NewOrderRepository(db)
	addressService  service.AddressService     = service.NewAddressService()
	orderService    service.OrderService       = service.NewOrderService(orderRepository, addressService)
	orderController controller.OrderController = controller.NewOrderController(orderService, addressService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	orderRoutes := r.Group("/api/order")
	{
		orderRoutes.GET("/", orderController.GetAllOrders)
		orderRoutes.POST("/", orderController.InsertOrder)
		orderRoutes.GET("/:id", orderController.FindOrderByID)
		orderRoutes.PUT("/:id", orderController.UpdateOrder)
		orderRoutes.DELETE("/:id", orderController.DeleteOrder)
	}
	r.Run(":2222")
}
