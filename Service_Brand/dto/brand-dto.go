package dto

type BrandUpdateDTO struct {
    ID    uint   `json:"id" form:"id" binding:"required"`
    Name  string `json:"name" form:"name" binding:"required,min=3,max=255"`
    Image string `json:"image" form:"image" binding:"omitempty"`
}

type BrandCreateDTO struct {
    Name  string `json:"name" form:"name" binding:"required,min=3,max=255"`
    Image string `json:"image" form:"image" binding:"omitempty"`
}
