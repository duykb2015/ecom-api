package usecase

import (
	"errors"
	"net/http"
	"strconv"

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
	user, err := u.r.GetByEmail(email)
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
	user, err := u.r.GetByEmail(request.Email)
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
		"name":     request.Name,
		"address":  request.Address,
		"phone":    request.Phone,
	}

	err = u.r.Create(insertData)

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

func (u *UserUsecase) GetInfo(tokenString string) (entity.UserInfoResponse, error) {
	user, err := u.r.GetInfo(tokenString)
	if err != nil {
		return entity.UserInfoResponse{}, err
	}

	resp := entity.UserInfoResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Address:  user.Address,
		Phone:    user.Phone,
		Status:   user.Status,
	}

	return resp, nil
}

func (u *UserUsecase) UpdateInfo(request entity.AuthRequest) error {
	userID, err := strconv.Atoi(request.ID)
	if err != nil {
		return err
	}
	user, err := u.r.GetByID(userID)
	if err != nil {
		return err
	}

	updateData := map[string]interface{}{
		"name":    request.Name,
		"address": request.Address,
		"phone":   request.Phone,
	}

	if request.OldPassword != "" {
		if user.Password != request.OldPassword {
			return errors.New("Old password dose not match")
		}

		updateData["password"] = request.NewPassword
	}

	err = u.r.Update(userID, updateData)
	return err
}

func (r *UserUsecase) RefreshToken() (interface{}, error) {
	return nil, nil
}
