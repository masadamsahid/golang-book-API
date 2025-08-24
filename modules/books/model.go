package books

import "time"

type Book struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	ImageURL    *string    `json:"image_url"`
	ReleaseYear int        `json:"release_year"`
	Price       *int       `json:"price"`
	TotalPage   *int       `json:"total_page"`
	Thickness   *string    `json:"thickness"`
	CategoryID  uint       `json:"category_id"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   string     `json:"created_by"`
	ModifiedAt  *time.Time `json:"modified_at"`
	ModifiedBy  *string    `json:"modified_by"`
}
