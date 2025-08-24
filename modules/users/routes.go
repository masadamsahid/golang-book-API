package users

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/register", HandleRegister)
	routerGroup.POST("/login", HandleLogin)
}
