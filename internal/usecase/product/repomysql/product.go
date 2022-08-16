package repomysql

import (
	"github.com/duykb2015/ecom-api/internal/entity"
	"gorm.io/gorm"
)

type ProductRepo struct {
	*gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db}
}

func (p *ProductRepo) GetAllProduct() ([]entity.Product, error) {
	product := []entity.Product{}
	p.Table("product").Find(&product)
	for i, val := range product {
		p.Table("product_items").Where("product_id", val.ID).Find(&product[i].ProductItems)
		for j, item := range product[i].ProductItems {
			p.Table("product_item_colors").Where("product_item_id", item.ID).Find(&product[i].ProductItems[j].ProductItemColors)
			p.Table("product_item_images").Where("product_item_id", item.ID).Find(&product[i].ProductItems[j].ProductItemImages)
		}
	}
	return product, nil
}
