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

func (p *ProductRepo) GetAll() ([]entity.Product, error) {
	product := []entity.Product{}
	err := p.Table("product").Where("status > ?", 0).Find(&product).Error
	return product, err
}

func (p *ProductRepo) GetByCategory(categoryIDs int) ([]entity.Product, error) {
	product := []entity.Product{}
	err := p.Table("product").Where("category_id = ? AND status > 0", categoryIDs).Find(&product).Error
	return product, err
}

func (p *ProductRepo) Get(productID int) (entity.Product, error) {
	productItems := entity.Product{}
	err := p.Table("product").Where("id = ?", productID).Find(&productItems).Error
	return productItems, err
}

func (p *ProductRepo) GetItems(itemID int) (entity.ProductItems, error) {
	productItems := entity.ProductItems{}
	err := p.Table("product_items").Where("id = ?", itemID).Find(&productItems).Error
	return productItems, err
}

func (p *ProductRepo) _GetAttributes(productID int, null bool) ([]entity.ProductAttributes, error) {
	attribute := []entity.ProductAttributes{}
	isNull := ""
	if null {
		isNull = "IS NULL"
	}
	err := p.Table("product_attributes").
		Joins("product_attribute_values").
		Where("product_id = ? AND product_item_id ?", isNull, productID).
		Find(&attribute).Error
	return attribute, err
}

func (p *ProductRepo) GetAttributes(productID int) ([]entity.ProductAttributes, error) {
	attribute := []entity.ProductAttributes{}
	err := p.Table("product_attributes pa").
		Select("pa.id, pa.product_id, pa.product_item_id, pav.name, pav.key, pav.value, pa.status, pa.created_at, pa.updated_at").
		Joins("JOIN product_attribute_values pav ON pa.product_attribute_value_id = pav.id").
		Where("product_id = ? AND product_item_id IS NULL", productID).
		Find(&attribute).Error
	return attribute, err
}

func (p *ProductRepo) GetItemAttributes(productItemID int) ([]entity.ProductAttributes, error) {
	attribute := []entity.ProductAttributes{}
	err := p.Table("product_attributes pa").
		Select("pa.id, pa.product_id, pa.product_item_id, pav.name, pav.key, pav.value, pa.status, pa.created_at, pa.updated_at").
		Joins("JOIN product_attribute_values pav ON pa.product_attribute_value_id = pav.id").
		Where("product_item_id = ?", productItemID).
		Find(&attribute).Error
	return attribute, err
}

func (p *ProductRepo) GetItemImages(itemId int) ([]entity.ProductItemImages, error) {
	attribute := []entity.ProductItemImages{}
	err := p.Table("product_item_images").
		Where("product_item_id = ?", itemId).
		Find(&attribute).Error
	return attribute, err
}

func (p *ProductRepo) GetItemColors(itemId int) ([]entity.ProductItemColors, error) {
	attribute := []entity.ProductItemColors{}
	err := p.Table("product_item_colors").
		Where("product_item_id = ?", itemId).
		Find(&attribute).Error
	return attribute, err
}
