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

	//Generate token
	tokenString, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return entity.AuthResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}, err

	}

	//Save token
	err = u.r.SaveToken(user.ID, tokenString)
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

func (u *UserUsecase) AuthRegister(request entity.AuthRequest) (entity.AuthResponse, error) {
	user, err := u.r.GetUser(request.Email)
	if err != nil {
		return entity.AuthResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		}, err
	}
	//Check user name and password
	if user.Email != "" {
		return entity.AuthResponse{
			Code:    http.StatusUnauthorized,
			Message: "Email already exist!",
		}, nil
	}

	insertData := map[string]interface{}{
		"email":    request.Email,
		"password": request.Password,
		"username": request.UserName,
		"address":  request.Address,
		"phone":    request.Phone,
	}

	err = u.r.CreateUser(insertData)

	if err != nil {
		return entity.AuthResponse{
			Code:    http.StatusInternalServerError,
			Message: "An error occurred, please try again later!",
		}, nil
	}

	return entity.AuthResponse{
		Code:    http.StatusOK,
		Message: "Ok!",
	}, nil
}

func (r *UserUsecase) RefreshToken() (interface{}, error) {
	return nil, nil
}
