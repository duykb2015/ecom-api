package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type CartRepo interface {
	GetByUserID(userID int) ([]entity.Cart, error)
}

type Cart interface {
	GetByUserID(userID int) ([]entity.CartResponse, error)
}
