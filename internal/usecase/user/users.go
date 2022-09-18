package usecase

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/duykb2015/ecom-api/internal/entity"
)

type UserUsecase struct {
	r UserRepo
}

// const secretKey = os.Getenv("SECRET_KEY")

func New(r UserRepo) *UserUsecase {
	return &UserUsecase{r}
}

func (r *UserUsecase) AuthLogin() (entity.UserResponse, error) {
	tokenString, err := generateToken("1")
	if err != nil {
		return entity.UserResponse{}, err
	}
	result := entity.UserResponse{
		ID:    1,
		Token: tokenString,
	}
	return result, nil
}

func generateToken(userID string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	})
	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}
