package categories

import (
	"database/sql"
	"log"
	"net/http"
	"quiz3/database/db"
	"quiz3/helpers"
	"quiz3/modules/books"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleCreateCategory(ctx *gin.Context) {
	u, ok := ctx.Get("user")
	if !ok {
		log.Println("Failed get user from context")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user, ok := u.(helpers.AuthPayload)
	if !ok {
		log.Println("Failed convert user from context to AuthPayload")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var createCategoryDto CreateCategoryDto
	err := ctx.ShouldBind(&createCategoryDto)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		log.Printf("%+v\n", validationErrors)
		errs := helpers.HandleValidationErrors(validationErrors)

		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed", "errors": errs})
		return
	}

	var newCategory Category
	sqlCreteNewCategory := `INSERT INTO categories (name, created_by) VALUES ($1, $2) RETURNING id, name, created_at, created_by , modified_at, modified_by`

	err = db.DBconn.QueryRow(sqlCreteNewCategory, createCategoryDto.Name, user.Username).Scan(
		&newCategory.ID,
		&newCategory.Name,
		&newCategory.CreatedAt,
		&newCategory.CreatedBy,
		&newCategory.ModifiedAt,
		&newCategory.ModifiedBy,
	)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed creating category",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success creating category",
		"data":    newCategory,
	})
}

func HandleGetAllCategories(ctx *gin.Context) {
	sqlGetAllGetCategories := `SELECT id, name, created_at, created_by , modified_at, modified_by FROM categories ORDER BY created_at ASC`
	rows, err := db.DBconn.Query(sqlGetAllGetCategories)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed retrieving categories",
		})
		return
	}

	var categories []Category

	defer rows.Close()
	for rows.Next() {
		var c Category
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.CreatedAt,
			&c.CreatedBy,
			&c.ModifiedAt,
			&c.ModifiedBy,
		)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed retrieving categories",
			})
			return
		}
		categories = append(categories, c)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success retrieved categories",
		"data":    categories,
	})
}

func HandleGetCategoryByID(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid category ID",
		})
		return
	}

	var category Category
	sqlGetCategoryByID := `SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories WHERE id = $1`
	err = db.DBconn.QueryRow(sqlGetCategoryByID, id).Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.CreatedBy,
		&category.ModifiedAt,
		&category.ModifiedBy,
	)
	if err != nil {

		status := http.StatusInternalServerError
		msg := "Failed retrieving categories"

		if err.Error() == sql.ErrNoRows.Error() {
			status = http.StatusNotFound
			msg = "Category not found"
		}

		log.Println(err)
		ctx.JSON(status, gin.H{
			"message": msg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success retrieved category",
		"data":    category,
	})
}

func HandleUpdateCategoryByID(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid category ID",
		})
		return
	}

	u, ok := ctx.Get("user")
	if !ok {
		log.Println("Failed get user from context")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user, ok := u.(helpers.AuthPayload)
	if !ok {
		log.Println("Failed convert user from context to AuthPayload")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var updateCategoryDto UpdateCategoryDto
	err = ctx.ShouldBind(&updateCategoryDto)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		log.Printf("%+v\n", validationErrors)
		errs := helpers.HandleValidationErrors(validationErrors)

		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed", "errors": errs})
		return
	}

	var updatedCategory Category
	sqlUpdateCategory := `UPDATE categories SET name = $2, modified_by = $3, modified_at = NOW() WHERE id = $1 RETURNING id, name, created_at, created_by, modified_at, modified_by`
	err = db.DBconn.QueryRow(sqlUpdateCategory, id, updateCategoryDto.Name, user.Username).Scan(
		&updatedCategory.ID,
		&updatedCategory.Name,
		&updatedCategory.CreatedAt,
		&updatedCategory.CreatedBy,
		&updatedCategory.ModifiedAt,
		&updatedCategory.ModifiedBy,
	)
	if err != nil {
		status := http.StatusInternalServerError
		msg := "Failed updating category"

		if err.Error() == sql.ErrNoRows.Error() {
			status = http.StatusNotFound
			msg = "Category not found"
		}

		log.Println(err)
		ctx.JSON(status, gin.H{
			"message": msg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success updating category",
		"data":    updatedCategory,
	})
}

func HandleDeleteCategoryByID(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid category ID",
		})
		return
	}

	sqlDeleteCategory := `DELETE FROM categories WHERE id = $1 RETURNING id`
	var deletedID int
	err = db.DBconn.QueryRow(sqlDeleteCategory, id).Scan(&deletedID)
	if err != nil {
		status := http.StatusInternalServerError
		msg := "Failed deleting category"

		if err.Error() == sql.ErrNoRows.Error() {
			status = http.StatusNotFound
			msg = "Category not found"
		}

		log.Println(err)
		ctx.JSON(status, gin.H{
			"message": msg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success deleting category",
	})
}

func HandleGetAllBooksFromCategoryByID(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid category ID",
		})
		return
	}

	sqlGetAllBooksFromCategory := `
SELECT
	b.id,
	b.title,
	b.description,
	b.image_url,
	b.release_year,
	b.price,
	b.total_page,
	b.thickness,
	b.category_id,
	b.created_at,
	b.created_by,
	b.modified_at,
	b.modified_by
FROM
	books b
JOIN
	categories c ON b.category_id = c.id
WHERE
	b.category_id = $1
ORDER BY
	b.created_at ASC
`
	rows, err := db.DBconn.Query(sqlGetAllBooksFromCategory, id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed retrieving books from category",
		})
		return
	}

	var bookList []books.Book

	defer rows.Close()
	for rows.Next() {
		var b books.Book
		err := rows.Scan(
			&b.ID,
			&b.Title,
			&b.Description,
			&b.ImageURL,
			&b.ReleaseYear,
			&b.Price,
			&b.TotalPage,
			&b.Thickness,
			&b.CategoryID,
			&b.CreatedAt,
			&b.CreatedBy,
			&b.ModifiedAt,
			&b.ModifiedBy,
		)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed retrieving books from category",
			})
			return
		}
		bookList = append(bookList, b)
	}

	if len(bookList) < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "This category has no book yet",
			"data":    bookList,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success retrieved books from category",
		"data":    bookList,
	})
}
