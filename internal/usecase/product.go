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
