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

func (p *ProductRepo) GetAllProductLine() ([]entity.Product, error) {
	product := []entity.Product{}
	p.Table("product").Where("status > ?", 0).Find(&product)
	return product, nil
}

func (p *ProductRepo) GetAllProductLineByCategory(slug string) ([]entity.Product, error) {
	product := []entity.Product{}
	p.Table("product").Where("category_id = ? AND status > ?", slug, 0).Find(&product)
	return product, nil
}

func (p *ProductRepo) GetAllProductItemsByProductLine(id int) ([]entity.Product, error) {
	product := []entity.Product{}
	p.Table("product").Where("id = ? AND status > ?", id, 0).Find(&product)
	for i, val := range product {
		p.Table("product_items").Where("product_id", val.ID).Find(&product[i].ProductItems)
	}
	return product, nil
}

func (p *ProductRepo) GetProductItemBySlug(id int, slug string) (entity.Product, error) {
	product := entity.Product{}
	field := "product_attributes.id, product_id, product_item_id, pav.name, pav.key, pav.value, pav.created_at, pav.updated_at"
	clause := "JOIN product_attribute_values pav ON pav.id = product_attributes.product_attribute_value_id"

	p.Table("product").Where("id", id).Find(&product)
	p.Table("product_attributes").
		Select(field).
		Joins(clause).
		Where("product_id", id).
		Find(&product.ProductAttributes)
	p.Table("product_items").Where("product_id", product.ID).Find(&product.ProductItems)
	for i, val := range product.ProductItems {
		p.Table("product_attributes").
			Select(field).
			Joins(clause).
			Where("product_item_id", val.ID).
			Find(&product.ProductItems[i].ProductAttributes)
		p.Table("product_item_colors").Where("product_item_id", val.ID).Find(&product.ProductItems[i].ProductItemColors)
		p.Table("product_item_images").Where("product_item_id", val.ID).Find(&product.ProductItems[i].ProductItemImages)
	}
	return product, nil
}
