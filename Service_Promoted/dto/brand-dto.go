package dto

type PromotedUpdateDTO struct {
	ID    uint   `json:"id" form:"id"`
	Title string `json:"title" form:"title" binding:"required,min=3,max=255"`
	Image string `json:"image" form:"image" binding:"omitempty"`
}

type PromotedCreateDTO struct {
	Title string `json:"title" form:"title" binding:"required,min=3,max=255"`
	Image string `json:"image" form:"image" binding:"omitempty"`
}
