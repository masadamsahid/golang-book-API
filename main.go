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

	scalargo "github.com/bdpiprava/scalar-go"
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

	router.GET("/api/hello-world", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello, world! Server is running",
		})
	})

	api := router.Group("/api")

	users.UserRoutes(api.Group("/users"))
	categories.CategoriesRoutes(api.Group("/categories"))
	books.BooksRoutes(api.Group("/books"))

	// OpenAPI
	openAPIYAMLDocs, err := os.ReadFile("./docs/swagger.yaml")
	if err != nil {
		panic(err)
	}

	router.GET("/openapi", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/plain", openAPIYAMLDocs)
	})

	// Scalar

	html, err := scalargo.NewV2(
		scalargo.WithSpecURL("/openapi"),
		scalargo.WithTheme(scalargo.ThemeBluePlanet),
		scalargo.WithMetaDataOpts(
			scalargo.WithTitle("📚 Golang Book Management API"),
		),
	)
	router.GET("/docs", func(ctx *gin.Context) {
		if err != nil {
			panic(err)
		}
		ctx.Data(http.StatusOK, "text/html", []byte(html))
	})

	PORT := os.Getenv("APP_PORT")
	if PORT == "" {
		PORT = "8080"
	}

	router.Run(":" + PORT)
}
