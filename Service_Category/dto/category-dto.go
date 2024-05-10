package dto

type CategoryUpdateDTO struct {
    ID    uint   `json:"id" form:"id" binding:"required"`
    Name  string `json:"name" form:"name" binding:"required,min=3,max=255"`
}

type CategoryCreateDTO struct {
    Name  string `json:"name" form:"name" binding:"required,min=3,max=255"`
}
