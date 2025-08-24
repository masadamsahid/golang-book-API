package books

import (
	"database/sql"
	"log"
	"net/http"
	"quiz3/database/db"
	"quiz3/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleCreateBook(ctx *gin.Context) {
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

	var createBookDto CreateBookDto
	err := ctx.ShouldBind(&createBookDto)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		errs := helpers.HandleValidationErrors(validationErrors)

		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed", "errors": errs})
		return
	}

	var newBook Book
	sqlCreteNewBook := `INSERT INTO
books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_by)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by
`
	thickness := "tipis"
	if createBookDto.TotalPage >= 100 {
		thickness = "tebal"
	}

	err = db.DBconn.QueryRow(sqlCreteNewBook,
		createBookDto.Title,
		createBookDto.Description,
		createBookDto.ImageURL,
		createBookDto.ReleaseYear,
		createBookDto.Price,
		createBookDto.TotalPage,
		thickness,
		createBookDto.CategoryID,
		user.Username,
	).Scan(
		&newBook.ID,
		&newBook.Title,
		&newBook.Description,
		&newBook.ImageURL,
		&newBook.ReleaseYear,
		&newBook.Price,
		&newBook.TotalPage,
		&newBook.Thickness,
		&newBook.CategoryID,
		&newBook.CreatedAt,
		&newBook.CreatedBy,
		&newBook.ModifiedAt,
		&newBook.ModifiedBy,
	)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed creating new book",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success creating book",
		"data":    newBook,
	})
}

func HandleGetAllBooks(ctx *gin.Context) {
	sqlGetAllBooks := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books ORDER BY created_at ASC`
	rows, err := db.DBconn.Query(sqlGetAllBooks)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed retrieving books",
		})
		return
	}

	var books []Book

	defer rows.Close()
	for rows.Next() {
		var b Book
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
				"message": "Failed retrieving books",
			})
			return
		}
		books = append(books, b)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success retrieving books",
		"data":    books,
	})
}

func HandleGetBookByID(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid book ID",
		})
		return
	}

	var book Book
	sqlGetBookByID := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books WHERE id = $1 LIMIT 1`
	err = db.DBconn.QueryRow(sqlGetBookByID, id).Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.ImageURL,
		&book.ReleaseYear,
		&book.Price,
		&book.TotalPage,
		&book.Thickness,
		&book.CategoryID,
		&book.CreatedAt,
		&book.CreatedBy,
		&book.ModifiedAt,
		&book.ModifiedBy,
	)
	if err != nil {
		status := http.StatusInternalServerError
		msg := "Failed retrieving book"

		if err.Error() == sql.ErrNoRows.Error() {
			status = http.StatusNotFound
			msg = "Book not found"
		}

		log.Println(err)
		ctx.JSON(status, gin.H{
			"message": msg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success retrieving book",
		"data":    book,
	})
}

func HandleUpdateBookByID(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid book ID",
		})
		return
	}

	var updatedBook Book
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

	var updateBookDto UpdateBookDto
	err = ctx.ShouldBind(&updateBookDto)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		errs := helpers.HandleValidationErrors(validationErrors)

		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Validation failed", "errors": errs})
		return
	}

	log.Printf("%+v\n", updateBookDto)

	thickness := "tipis"
	if updateBookDto.TotalPage >= 100 {
		thickness = "tebal"
	}

	sqlUpdateBook := `UPDATE books SET title = $2, description = $3, image_url = $4, release_year = $5, price = $6, total_page = $7, thickness = $8, category_id = $9, modified_by = $10, modified_at = NOW() WHERE id = $1
RETURNING id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by`

	err = db.DBconn.QueryRow(sqlUpdateBook,
		id,
		updateBookDto.Title,
		updateBookDto.Description,
		updateBookDto.ImageURL,
		updateBookDto.ReleaseYear,
		updateBookDto.Price,
		updateBookDto.TotalPage,
		thickness,
		updateBookDto.CategoryID,
		user.Username,
	).Scan(
		&updatedBook.ID,
		&updatedBook.Title,
		&updatedBook.Description,
		&updatedBook.ImageURL,
		&updatedBook.ReleaseYear,
		&updatedBook.Price,
		&updatedBook.TotalPage,
		&updatedBook.Thickness,
		&updatedBook.CategoryID,
		&updatedBook.CreatedAt,
		&updatedBook.CreatedBy,
		&updatedBook.ModifiedAt,
		&updatedBook.ModifiedBy,
	)
	if err != nil {
		status := http.StatusInternalServerError
		msg := "Failed updating book"

		if err.Error() == sql.ErrNoRows.Error() {
			status = http.StatusNotFound
			msg = "Book not found"
		}

		log.Println(err)
		ctx.JSON(status, gin.H{
			"message": msg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success updating book",
		"data":    updatedBook,
	})
}

func HandleDeleteBookByID(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid book ID",
		})
		return
	}

	sqlDeleteBook := `DELETE FROM books WHERE id = $1`
	res, err := db.DBconn.Exec(sqlDeleteBook, id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed deleting book",
		})
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed deleting book",
		})
		return
	}

	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success deleting book",
	})
}
