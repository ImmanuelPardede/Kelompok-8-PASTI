package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iqbalsiagian17/Service_Cart/dto"
	"github.com/iqbalsiagian17/Service_Cart/service"
)

type CartController interface {
	All(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type cartController struct {
	CartService service.CartService
}

func NewCartController(cartService service.CartService) CartController {
	return &cartController{
		CartService: cartService,
	}
}

func (cc *cartController) All(ctx *gin.Context) {
	carts := cc.CartService.All()
	ctx.JSON(http.StatusOK, carts)
}

func (cc *cartController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	cartID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
		return
	}
	cart := cc.CartService.FindByID(uint64(cartID))
	if cart.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}
	ctx.JSON(http.StatusOK, cart)
}

func (cc *cartController) Insert(ctx *gin.Context) {
	var cart dto.CartCreateDTO
	if err := ctx.ShouldBindJSON(&cart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCart := cc.CartService.Insert(cart)
	ctx.JSON(http.StatusCreated, newCart)
}

func (cc *cartController) Update(ctx *gin.Context) {
	var cart dto.CartUpdateDTO
	if err := ctx.ShouldBindJSON(&cart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedCart := cc.CartService.Update(cart)
	ctx.JSON(http.StatusOK, updatedCart)
}

func (cc *cartController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	cartID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
		return
	}
	// Periksa apakah ada cart dengan ID yang valid sebelum menghapus
	cart := cc.CartService.FindByID(uint64(cartID))
	if cart.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}
	cc.CartService.Delete(cart) // Menghapus cart dengan menggunakan data cart yang ditemukan
	ctx.JSON(http.StatusOK, gin.H{"message": "Cart deleted successfully"})
}
