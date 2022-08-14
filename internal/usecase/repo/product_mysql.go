package repo

import (
	"github.com/duykb2015/login-api/internal/entity"
	"gorm.io/gorm"
)

type ProductRepo struct {
	*gorm.DB
}

func New(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db}
}

func (p *ProductRepo) GetAllProduct() ([]entity.Product, error) {
	entities := []entity.Product{}
	p.Table("product").Find(&entities)
	return entities, nil
}
