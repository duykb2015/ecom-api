package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type ProductUsecase struct {
	repo ProductRepo
}

func NewProduct(r ProductRepo) *ProductUsecase {
	return &ProductUsecase{r}
}

func (p *ProductUsecase) GetAllProductLine() ([]entity.Product, error) {
	return p.repo.GetAllProductLine()
}

func (p *ProductUsecase) GetAllProductLineByCategory(slug string) ([]entity.Product, error) {
	return p.repo.GetAllProductLineByCategory(slug)
}

func (p *ProductUsecase) GetAllProductItemsByProductLine(id int) ([]entity.Product, error) {
	return p.repo.GetAllProductItemsByProductLine(id)
}

func (p *ProductUsecase) GetProductItemBySlug(id int, slug string) (entity.Product, error) {
	return p.repo.GetProductItemBySlug(id, slug)
}
