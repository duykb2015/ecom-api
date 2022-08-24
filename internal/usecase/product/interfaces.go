package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type (
	// ProductUsecase interface
	ProductRepo interface {
		GetAllProduct() ([]entity.Product, error)
		GetAllProductByCategory(id int) ([]entity.Product, error)
		GetAllProductItemsByProduct(id int) ([]entity.Product, error)
		GetProductItemInfo(product_id int, product_item_id int) (entity.Product, error)
	}

	Product interface {
		GetAll() ([]entity.ProductLineRespond, error)
		GetByCategory(id int) ([]entity.ProductLineRespond, error)
		Items(id int) ([]entity.ProductRespond, error)
		ItemInfo(product_id int, product_item_id int) (entity.ProductRespond, error)
	}
)
