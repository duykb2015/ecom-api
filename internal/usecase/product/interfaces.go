package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type (
	// ProductUsecase interface
	ProductRepo interface {
		GetAllProductLine() ([]entity.Product, error)
		GetAllProductLineByCategory(id int) ([]entity.Product, error)
		GetAllProductItemsByProductLine(id int) ([]entity.Product, error)
		//id Product Line Id, slug Product Item Slug
		GetProductItemInfo(product_id int, product_item_id int) (entity.Product, error)
		//a
	}
)
