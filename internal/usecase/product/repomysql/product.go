package repomysql

import (
	"github.com/duykb2015/ecom-api/internal/entity"
	"gorm.io/gorm"
)

type ProductRepo struct {
	*gorm.DB
}

func New(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db}
}

func (p *ProductRepo) GetAllProduct() ([]entity.Product, error) {
	product := []entity.Product{}
	err := p.Table("product").Where("status > ?", 0).Find(&product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductRepo) GetAllProductByCategory(id int) ([]entity.Product, error) {
	product := []entity.Product{}
	err := p.Table("product").Select("id, category_id, name, slug, status").Where("category_id = ? AND status > ?", id, 0).Find(&product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductRepo) GetAllProductItemsByProduct(id int) ([]entity.Product, error) {
	product := []entity.Product{}
	err := p.Table("product").Select("id, category_id, name, slug, status").Where("id = ? AND status > ?", id, 0).Find(&product).Error
	if err != nil {
		return nil, err
	}

	for i, val := range product {
		err = p.Table("product_items").Select("id, product_id, name, slug, status").Where("product_id", val.ID).Find(&product[i].ProductItems).Error
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}

func (p *ProductRepo) GetProductItemInfo(product_id int, product_item_id int) (entity.Product, error) {
	product := entity.Product{}
	field := "product_attributes.id, product_id, product_item_id, pav.name, pav.key, pav.value"
	clause := "JOIN product_attribute_values pav ON pav.id = product_attributes.product_attribute_value_id"

	err := p.Table("product").Where("id", product_id).Find(&product).Error
	if err != nil {
		return entity.Product{}, err
	}

	err = p.Table("product_attributes").
		Select(field).
		Joins(clause).
		Where("product_id = ? AND product_item_id IS NULL", product_id).
		Find(&product.ProductAttributes).Error
	if err != nil {
		return entity.Product{}, err
	}

	err = p.Table("product_items").Where("id", product_item_id).Find(&product.ProductItems).Error
	if err != nil {
		return entity.Product{}, err
	}

	err = p.Table("product_attributes").
		Select(field).
		Joins(clause).
		Where("product_item_id", product_item_id).
		Find(&product.ProductItems[0].ProductAttributes).Error
	if err != nil {
		return entity.Product{}, err
	}

	err = p.Table("product_item_colors").Where("product_item_id", product_item_id).Find(&product.ProductItems[0].ProductItemColors).Error
	if err != nil {
		return entity.Product{}, err
	}

	err = p.Table("product_item_images").Where("product_item_id", product_item_id).Find(&product.ProductItems[0].ProductItemImages).Error
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}
