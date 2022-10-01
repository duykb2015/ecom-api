package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type UserRepo interface {
	GetByID(id int) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
	SaveToken(userID uint, tokenString string) error
	Create(insertData map[string]interface{}) error
	Update(userID int, updateData map[string]interface{}) error
	GetInfo(tokenString string) (entity.User, error)
}

type User interface {
	AuthLogin(email string, password string) (entity.AuthResponse, error)
	AuthRegister(request entity.AuthRequest) (entity.AuthResponse, error)
	GetInfo(tokenString string) (entity.UserInfoResponse, error)
	UpdateInfo(request entity.AuthRequest) error
}
