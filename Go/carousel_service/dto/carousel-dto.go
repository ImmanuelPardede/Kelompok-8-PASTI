package dto

// CarouselItemCreateDTO adalah DTO untuk membuat CarouselItem
type CarouselCreateDTO struct {
	ID         uint64 `json:"id" form:"id" binding:"required"`
	ImageURL   string `json:"image_url" binding:"required"`
	Caption    string `json:"caption"`
	Subcaption string `json:"subcaption"`
}

// CarouselItemUpdateDTO adalah DTO untuk memperbarui CarouselItem
type CarouselUpdateDTO struct {
	ID         uint64 `json:"id" form:"id" binding:"required"`
	ImageURL   string `json:"image_url"`
	Caption    string `json:"caption"`
	Subcaption string `json:"subcaption"`
}
