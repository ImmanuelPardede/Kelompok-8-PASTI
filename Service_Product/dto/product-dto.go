package dto

type ProductCreateDTO struct {
	Name          string `json:"name" binding:"required"`
	CategoryID    uint   `json:"category_id" binding:"required"`
	SubCategoryID uint   `json:"subcategory_id" binding:"required"`
	BrandID       uint   `json:"brand_id" binding:"required"`
	Size          string `json:"size" binding:"required"`
	Quantity      int    `json:"quantity" binding:"required"`
	Price         int    `json:"price" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Image         string `json:"image" binding:"required"` // Tambahan bidang untuk gambar
}

type ProductUpdateDTO struct {
	ID            uint   `json:"id" form:"id"`
	Name          string `json:"name"`
	CategoryID    uint   `json:"category_id"`
	SubCategoryID uint   `json:"subcategory_id"`
	BrandID       uint   `json:"brand_id"`
	Size          string `json:"size"`
	Quantity      int    `json:"quantity"`
	Price         int    `json:"price"`
	Description   string `json:"description"`
	Image         string `json:"image"` // Tambahan bidang untuk gambar
}
