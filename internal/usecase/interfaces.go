package usecase

import "github.com/duykb2015/login-api/internal/entity"

type (
	// ProductUsecase interface
	ProductRepo interface {
		GetAllProduct() ([]entity.Product, error)
		GetProductByID(id string) ([]entity.Product, error)
	}

	MenuRepo interface {
		GetAllMenu() ([]entity.Menu, error)
	}
)
