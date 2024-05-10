package dto

type SubCategoryUpdateDTO struct {
	ID         uint   `json:"id" form:"id" binding:"required"`
	Name       string `json:"name" form:"name" binding:"required,min=3,max=255"`
	CategoryID uint   `json:"category_id" form:"category_id" binding:"required"`
}

type SubCategoryCreateDTO struct {
	Name       string `json:"name" form:"name" binding:"required,min=3,max=255"`
	CategoryID uint   `json:"category_id" form:"category_id" binding:"required"`
}
