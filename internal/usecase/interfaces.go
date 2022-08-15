package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type (
	// ProductUsecase interface
	ProductRepo interface {
		GetAllProduct() ([]entity.Product, error)
	}

	MenuRepo interface {
		GetAllMenu() ([]entity.Menu, error)
	}
)
