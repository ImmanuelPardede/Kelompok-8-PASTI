package dto

type CartUpdateDTO struct {
	ID   uint   `json:"id" form:"id"`
	Name string `json:"name" form:"name" binding:"required,min=3,max=255"`
}

type CartCreateDTO struct {
	Name string `json:"name" form:"name" binding:"required,min=3,max=255"`
}
