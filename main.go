package main

import (
	"log"
	"net/http"
	"os"
	"quiz3/database/db"
	"quiz3/modules/books"
	"quiz3/modules/categories"
	"quiz3/modules/users"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// load .env to os lib
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Postgres
	db.ConnectPg()
	defer db.StopDBConn()

	// Gin
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello, world! Server is running",
		})
	})

	api := router.Group("/api")

	users.UserRoutes(api.Group("/users"))
	categories.CategoriesRoutes(api.Group("/categories"))
	books.BooksRoutes(api.Group("/books"))

	PORT := os.Getenv("APP_PORT")
	if PORT == "" {
		PORT = "8080"
	}

	router.Run(":" + PORT)
}
