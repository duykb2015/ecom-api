package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type UserRepo interface {
	GetUser(email string) (entity.User, error)
	SaveToken(userID uint, tokenString string) error
	CreateUser(map[string]interface{}) error
}

type User interface {
	AuthLogin(email string, password string) (entity.AuthResponse, error)
	AuthRegister(request entity.AuthRequest) (entity.AuthResponse, error)
}
