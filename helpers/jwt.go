package helpers

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// var randJWTSecret = rand.Text()
// var randJWTSecretArrOfByte = []byte(randJWTSecret)

var jwtSecret = "*&(HD!)&EO"
var jwtSecretArrOfByte = []byte(jwtSecret)

type AuthTokenClaims struct {
	ID       int
	Username string
}

func CreateAuthToken(claims AuthTokenClaims) (string, error) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       claims.ID,
		"username": claims.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	log.Println("Token claim", tokenClaims)

	authToken, err := tokenClaims.SignedString(jwtSecretArrOfByte)
	if err != nil {
		log.Println("Error creating auth token:", err)
		return "", err
	}

	return authToken, nil
}
