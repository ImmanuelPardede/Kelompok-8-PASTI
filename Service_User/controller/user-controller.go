package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iqbalsiagian17/User/dto"
	"github.com/iqbalsiagian17/User/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService}
}

func (uc *UserController) Register(c *gin.Context) {
	var userDTO dto.UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate userDTO (optional)

	newUser, err := uc.userService.Register(userDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newUser})
}

func (uc *UserController) Login(c *gin.Context) {
	var loginDTO dto.LoginDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate loginDTO (optional)

	token, err := uc.userService.Login(loginDTO)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) Logout(c *gin.Context) {
	// Mengambil token dari header Authorization
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization token"})
		return
	}

	// Memanggil metode Logout dari service
	if err := uc.userService.Logout(tokenString); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Memberikan tanggapan sukses
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
