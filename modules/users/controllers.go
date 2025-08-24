package users

import (
	"log"
	"net/http"
	"quiz3/database/db"
	"quiz3/helpers"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func HandleRegister(ctx *gin.Context) {
	var newUserDto RegisteUserDto

	err := ctx.ShouldBind(&newUserDto)
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

	hashedPwd, err := helpers.HashPassword(newUserDto.Password)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	var newUser User

	sqlCreateNewUser := `INSERT INTO users (username, "password") VALUES ($1, $2) RETURNING id, username`

	err = db.DBconn.QueryRow(sqlCreateNewUser, newUserDto.Username, hashedPwd).Scan(
		&newUser.ID,
		&newUser.Username,
	)
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "unique constraint") {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": "'username' already taken",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed registering new user",
		})
		return
	}

	authToken, err := helpers.CreateAuthToken(helpers.AuthTokenClaims{
		ID:       newUser.ID,
		Username: newUser.Username,
	})
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed creating token",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success registering user",
		"data":    authToken,
	})
}

func HandleLogin(ctx *gin.Context) {
	var loginDto LoginUserDto

	err := ctx.ShouldBind(&loginDto)
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

	var user User

	sqlGetUserByUsername := `SELECT id, username, "password" FROM users WHERE username = $1`

	err = db.DBconn.QueryRow(sqlGetUserByUsername, loginDto.Username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
	)
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "no rows in result set") {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed logging in",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password))
	if err != nil {
		if err.Error() == bcrypt.ErrMismatchedHashAndPassword.Error() {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Wrong credentials",
			})
			return
		}

		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	authToken, err := helpers.CreateAuthToken(helpers.AuthTokenClaims{
		ID:       user.ID,
		Username: user.Username,
	})
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed creating token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success logging in",
		"data":    authToken,
	})
}
