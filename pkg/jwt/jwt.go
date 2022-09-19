package jwt

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = os.Getenv("SECRET_KEY")

func GenerateToken(userID string) (string, error) {
	var secretKey = os.Getenv("SECRET_KEY")
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"exp":        time.Now().Add(time.Minute * 15).Unix(),
	})
	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
