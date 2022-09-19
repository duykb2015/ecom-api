package usecase

import (
	"github.com/duykb2015/ecom-api/internal/entity"
)

type UserUsecase struct {
	r UserRepo
}

func New(r UserRepo) *UserUsecase {
	return &UserUsecase{r}
}

func (r *UserUsecase) AuthLogin() (entity.AuthResponse, error) {
	tokenString, err := generateToken("1")
	if err != nil {
		return entity.AuthResponse{}, err
	}
	result := entity.AuthResponse{
		Token: tokenString,
	}
	return result, nil
}

func (r *UserUsecase) AuthRegister() (entity.AuthResponse, error) {
	tokenString, err := GenerateToken("1")
	if err != nil {
		return entity.AuthResponse{}, err
	}
	result := entity.AuthResponse{
		Token: tokenString,
	}
	return result, nil
}

func (r *UserUsecase) RefreshToken() (interface{}, error) {
	return nil, nil
}
