package users

import (
	"net/http"
	"quiz3/helpers"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// fmt.Println("Form JWT middleware")
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		strAuthToken := parts[1]
		authToken, err := helpers.VerifyAuthToken(strAuthToken)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims, ok := authToken.Claims.(jwt.MapClaims); ok && authToken.Valid {
			ctx.Set("user", helpers.AuthPayload{
				ID:       int(claims["id"].(float64)),
				Username: claims["username"].(string),
				Exp:      int(claims["exp"].(float64)),
			})
		}

		ctx.Next()
	}
}
