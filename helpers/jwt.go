package helpers

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// var randJWTSecret = rand.Text()
// var randJWTSecretArrOfByte = []byte(randJWTSecret)

var jwtSecret = "*&(HD!)&EO"
var jwtSecretArrOfByte = []byte(jwtSecret)

type AuthPayload struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Exp      int    `json:"exp"`
}

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

func VerifyAuthToken(strAuthToken string) (*jwt.Token, error) {
	authToken, err := jwt.Parse(strAuthToken, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return jwtSecretArrOfByte, nil
	})

	// log.Println(err)

	return authToken, err
}
