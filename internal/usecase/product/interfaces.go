package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type (
	// ProductUsecase interface
	ProductRepo interface {
		GetAll() ([]entity.Product, error)
		GetByCategory(categoryID int) ([]entity.Product, error)
		GetItems(productID int) ([]entity.ProductItems, error)
		// GetItemInfo(productID int, productItemID int) (entity.Product, error)
	}

	Product interface {
		GetAll() ([]entity.ProductLineResponse, error)
		ByCategory(categoryID int) ([]entity.ProductLineResponse, error)
		Items(ProductID int) ([]entity.ProductResponse, error)
	}
)
