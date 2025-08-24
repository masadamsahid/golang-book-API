package categories

type CreateCategoryDto struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategoryDto struct {
	Name string `json:"name"`
}
