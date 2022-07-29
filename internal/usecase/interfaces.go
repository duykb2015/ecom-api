package usecase

import "github.com/duykb2015/login-api/internal/entity"

type (
	// ProductUsecase interface
	Product interface {
		GetAllProduct() ([]entity.Product, error)
		GetProductByID(id string) ([]entity.Product, error)
		CreateProduct(product Product) ([]entity.Product, error)
		UpdateProduct(product Product) ([]entity.Product, error)
		DeleteProduct(id string) error
	}

	ProductRepo interface {
		GetAllProduct() ([]entity.Product, error)
	}
)
