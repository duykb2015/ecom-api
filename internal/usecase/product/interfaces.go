package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type (
	// ProductUsecase interface
	ProductRepo interface {
		GetAllProductLine() ([]entity.Product, error)
		GetAllProductLineByCategory(slug string) ([]entity.Product, error)
		GetAllProductItemsByProductLine(id int) ([]entity.Product, error)
		//id Product Line Id, slug Product Item Slug
		GetProductItemBySlug(id int, slug string) (entity.Product, error)
	}
)
