package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iqbalsiagian17/User/config"
	"github.com/iqbalsiagian17/User/controller"
	"github.com/iqbalsiagian17/User/repository"
	"github.com/iqbalsiagian17/User/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = config.SetupDatabaseConnection()
	userRepository *repository.UserRepository = repository.NewUserRepository(db)
	userService    service.UserService        = *service.NewUserService(userRepository)
	userController controller.UserController  = *controller.NewUserController(&userService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()


	
	/* // Periksa apakah sudah ada admin
	var adminCount int64
	if err := db.Model(&model.User{}).Where("role = ?", "admin").Count(&adminCount).Error; err != nil {
		log.Fatalf("Error checking admin count: %v", err)
	}

	// Jika belum ada admin, buat data dummy untuk admin
	if adminCount == 0 {
		admin := model.User{
			Name:     "Admin",
			Email:    "admin@gmail.com",
			Password: "adminpassword",
			Role:     "admin",
		}

		if err := db.Create(&admin).Error; err != nil {
			log.Fatalf("Error creating admin: %v", err)
		}
		fmt.Println("Admin dummy created successfully")
	} else {
		fmt.Println("Admin already exists, skipping dummy creation")
	} */

	
	userRoutes := r.Group("/api/user")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
		userRoutes.POST("/logout", userController.Logout)
	}
	r.Run(":1111")

}
