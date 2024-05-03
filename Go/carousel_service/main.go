package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iqbalsiagian17/carousel_service/config"
	"github.com/iqbalsiagian17/carousel_service/controller"
	"github.com/iqbalsiagian17/carousel_service/repository"
	"github.com/iqbalsiagian17/carousel_service/service"
	"gorm.io/gorm"
)

var (
	db                 *gorm.DB                      = config.SetupDatabaseConnection()
	carouselRepository repository.CarouselRepository = repository.NewCarouselRepository(db)
	carouselService    service.CarouselService       = service.NewCarouselService(carouselRepository)
	carouselController controller.CarouselController = controller.NewCarouselController(carouselService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	carouselRoutes := r.Group("/api/carousel")
	{
		carouselRoutes.GET("/", carouselController.Index)
		carouselRoutes.POST("/create", carouselController.Create)
		carouselRoutes.GET("/:id", carouselController.Show)
		carouselRoutes.PUT("/:id", carouselController.Update)
		carouselRoutes.DELETE("/:id", carouselController.Delete)
	}

	r.Run(":8081")
}
