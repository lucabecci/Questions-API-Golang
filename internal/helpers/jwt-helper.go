package helpers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func JwtGenerator(userID uint) (string, error) {
	jwtSecret := os.Getenv("JWT_KEY")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = "14"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 1) //one day

	tkn, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tkn, nil
}
