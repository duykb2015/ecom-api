package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type UserRepo interface {
}

type User interface {
	AuthLogin() (entity.AuthResponse, error)
	AuthRegister() (entity.AuthResponse, error)
}
