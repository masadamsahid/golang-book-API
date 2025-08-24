package categories

import (
	"quiz3/modules/users"

	"github.com/gin-gonic/gin"
)

func CategoriesRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/", users.JwtAuthMiddleware(), HandleCreateCategory)
	routerGroup.GET("/", HandleGetAllCategories)
	routerGroup.GET("/:id", HandleGetCategoryByID)
	routerGroup.PUT("/:id", users.JwtAuthMiddleware(), HandleUpdateCategoryByID)
	routerGroup.DELETE("/:id", users.JwtAuthMiddleware(), HandleDeleteCategoryByID)
}
