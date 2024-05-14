package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iqbalsiagian17/Service_Cart/config"
	"github.com/iqbalsiagian17/Service_Cart/controller"
	"github.com/iqbalsiagian17/Service_Cart/repository"
	"github.com/iqbalsiagian17/Service_Cart/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	cartRepository repository.CartRepository = repository.NewCartRepository(db)
	cartService    service.CartService       = service.NewCartService(cartRepository)
	cartController controller.CartController = controller.NewCartController(cartService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	cartRoutes := r.Group("/api/cart")
	{
		cartRoutes.GET("/", cartController.All)
		cartRoutes.POST("/", cartController.Insert)
		cartRoutes.GET("/:id", cartController.FindByID)
		cartRoutes.PUT("/:id", cartController.Update)
		cartRoutes.DELETE("/:id", cartController.Delete)
	}
	r.Run(":1111")
}
