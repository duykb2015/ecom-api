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

func (p *ProductUsecase) GetAllProductLineByCategory(id int) ([]entity.Product, error) {
	return p.repo.GetAllProductLineByCategory(id)
}

func (p *ProductUsecase) GetAllProductItemsByProductLine(id int) ([]entity.Product, error) {
	return p.repo.GetAllProductItemsByProductLine(id)
}

func (p *ProductUsecase) GetProductItemInfo(product_id int, product_item_id int) (entity.Product, error) {
	return p.repo.GetProductItemInfo(product_id, product_item_id)
}
