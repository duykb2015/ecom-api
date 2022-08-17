package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type ProductUsecase struct {
	repo ProductRepo
}

func NewProduct(r ProductRepo) *ProductUsecase {
	return &ProductUsecase{r}
}

func (p *ProductUsecase) GetAllProduct() ([]entity.Product, error) {
	return p.repo.GetAllProduct()
}

func (p *ProductUsecase) GetAllProductByCategory(slug string) ([]entity.Product, error) {
	return p.repo.GetAllProductByCategory(slug)
}

func (p *ProductUsecase) GetAllProductByProductLine(slug string) ([]entity.Product, error) {
	return p.repo.GetAllProductByProductLine(slug)
}

func (p *ProductUsecase) GetProductBySlug(slug string) (entity.Product, error) {
	return p.repo.GetProductBySlug(slug)
}
