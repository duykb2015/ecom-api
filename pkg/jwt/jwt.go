package jwt

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func GenerateToken(userID uint) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("app - main - Load env error: %s", err)
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"exp":        time.Now().Add(time.Minute * 60).Unix(),
	})
	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}
