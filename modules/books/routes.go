package books

import (
	"quiz3/modules/users"

	"github.com/gin-gonic/gin"
)

func BooksRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/", users.JwtAuthMiddleware(), HandleCreateBook)
	routerGroup.GET("/", HandleGetAllBooks)
	routerGroup.GET("/:id", HandleGetBookByID)
	routerGroup.PUT("/:id", users.JwtAuthMiddleware(), HandleUpdateBookByID)
	routerGroup.DELETE("/:id", users.JwtAuthMiddleware(), HandleDeleteBookByID)
}
