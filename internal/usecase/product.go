package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type ProductUsecase struct {
	repo ProductRepo
}

func New(r ProductRepo) *ProductUsecase {
	return &ProductUsecase{r}
}

func (p *ProductUsecase) Product() ([]entity.Product, error) {
	return p.repo.GetAllProduct()
}
