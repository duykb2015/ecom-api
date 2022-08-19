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

func (p *ProductRepo) GetAllProductLineByCategory(id int) ([]entity.Product, error) {
	product := []entity.Product{}
	p.Table("product").Select("id, category_id, name, slug, status").Where("category_id = ? AND status > ?", id, 0).Find(&product)
	return product, nil
	//a
}

func (p *ProductRepo) GetAllProductItemsByProductLine(id int) ([]entity.Product, error) {
	product := []entity.Product{}
	p.Table("product").Select("id, category_id, name, slug, status").Where("id = ? AND status > ?", id, 0).Find(&product)
	for i, val := range product {
		p.Table("product_items").Select("id, product_id, name, slug, status").Where("product_id", val.ID).Find(&product[i].ProductItems)
	}
	return product, nil
}

func (p *ProductRepo) GetProductItemInfo(product_id int, product_item_id int) (entity.Product, error) {
	product := entity.Product{}
	field := "product_attributes.id, product_id, product_item_id, pav.name, pav.key, pav.value"
	clause := "JOIN product_attribute_values pav ON pav.id = product_attributes.product_attribute_value_id"

	p.Table("product").Where("id", product_id).Find(&product)
	p.Table("product_attributes").
		Select(field).
		Joins(clause).
		Where("product_id", product_id).
		Find(&product.ProductAttributes)

	p.Table("product_items").Where("product_id", product_item_id).Find(&product.ProductItems)
	p.Table("product_attributes").
		Select(field).
		Joins(clause).
		Where("product_item_id", product_item_id).
		Find(&product.ProductItems[0].ProductAttributes)
	p.Table("product_item_colors").Where("product_item_id", product_item_id).Find(&product.ProductItems[0].ProductItemColors)
	p.Table("product_item_images").Where("product_item_id", product_item_id).Find(&product.ProductItems[0].ProductItemImages)

	return product, nil
}
