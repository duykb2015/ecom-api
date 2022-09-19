package usecase

import (
	"net/http"

	"github.com/duykb2015/ecom-api/internal/entity"
	"github.com/duykb2015/ecom-api/pkg/jwt"
)

type UserUsecase struct {
	r UserRepo
}

func New(r UserRepo) *UserUsecase {
	return &UserUsecase{r}
}

func (u *UserUsecase) AuthLogin(email string, password string) (entity.AuthResponse, error) {
	user, err := u.r.GetUser(email)
	if err != nil {
		return entity.AuthResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		}, err
	}
	//Check user name and password
	if user.Email == "" || user.Password != password {
		return entity.AuthResponse{
			Code:    http.StatusUnauthorized,
			Message: "Wrong infomation, please check again",
		}, nil
	}

	tokenString, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return entity.AuthResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}, err
	}
	return entity.AuthResponse{
		Code:    http.StatusOK,
		Message: "Ok!",
		Token:   tokenString,
	}, nil
}

func (r *UserUsecase) AuthRegister() (entity.AuthResponse, error) {
	tokenString, err := jwt.GenerateToken(2)
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
