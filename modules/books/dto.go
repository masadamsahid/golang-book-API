package books

type CreateBookDto struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`
	ImageURL    *string `json:"image_url" binding:"omitempty,url"`
	ReleaseYear int     `json:"release_year" binding:"required"`
	Price       *int    `json:"price" binding:"omitempty,gte=0"`
	TotalPage   uint    `json:"total_page" binding:"omitempty,gte=1"`
	CategoryID  uint    `json:"category_id" binding:"required"`
}

type UpdateBookDto struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`
	ImageURL    *string `json:"image_url" binding:"omitempty,url"`
	ReleaseYear int     `json:"release_year" binding:"required"`
	Price       *int    `json:"price" binding:"omitempty,gte=0"`
	TotalPage   uint    `json:"total_page" binding:"required,omitempty,gte=1"`
	CategoryID  uint    `json:"category_id" binding:"required"`
}
