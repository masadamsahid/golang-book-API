package main

import (
	"quiz3/database/db"
	"quiz3/modules/books"
	"quiz3/modules/categories"
	"quiz3/modules/users"

	"github.com/gin-gonic/gin"
)

func main() {
	// Postgres
	db.ConnectPg()
	defer db.StopDBConn()

	// Gin
	router := gin.Default()

	api := router.Group("/api")

	users.UserRoutes(api.Group("/users"))
	categories.CategoriesRoutes(api.Group("/categories"))
	books.BooksRoutes(api.Group("/books"))

	router.Run()
}
