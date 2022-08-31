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
	err := p.Table("product").Where("category_id IN ? AND status > 0", categoryIDs).Find(&product).Error
	return product, err
}

func (p *ProductRepo) GetItems(productID int) ([]entity.ProductItems, error) {
	productItems := []entity.ProductItems{}
	err := p.Table("product_items").Where("product_id", productID).Find(&productItems).Error
	return productItems, err
}

func (p *ProductRepo) GetAttributes(productID int) ([]entity.ProductAttributes, error) {
	attribute := []entity.ProductAttributes{}
	err := p.Table("product_attributes").
		Where("product_id = ? AND product_item_id IS NULL", productID).
		Find(&attribute).Error
	return attribute, err
}

func (p *ProductRepo) GetItemAttributes(productItemID int) ([]entity.ProductAttributes, error) {
	attribute := []entity.ProductAttributes{}
	err := p.Table("product_attributes").
		Where("product_item_id = ?", productItemID).
		Find(&attribute).Error
	return attribute, err
}

func (p *ProductRepo) GetItemImages(itemId int) ([]entity.ProductAttributes, error) {
	attribute := []entity.ProductAttributes{}
	err := p.Table("product_item_colors").
		Where("product_item_id = ?", itemId).
		Find(&attribute).Error
	return attribute, err
}

func (p *ProductRepo) GetItemColors(itemId int) ([]entity.ProductAttributes, error) {
	attribute := []entity.ProductAttributes{}
	err := p.Table("product_item_images").
		Where("product_item_id = ?", itemId).
		Find(&attribute).Error
	return attribute, err
}
