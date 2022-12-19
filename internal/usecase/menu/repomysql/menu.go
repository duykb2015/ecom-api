package repomysql

import (
	"github.com/duykb2015/ecom-api/internal/entity"
	"gorm.io/gorm"
)

type MenuRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *MenuRepo {
	return &MenuRepo{db}
}

func (m *MenuRepo) GetParents() ([]entity.Menu, error) {
	menus := []entity.Menu{}
	err := m.db.Table("menu").
		Where("parent_id = 0 AND status > 0").
		Find(&menus).Error
	return menus, err
}

func (m *MenuRepo) GetChildens(menuIDs []int) ([]entity.Menu, error) {
	menus := []entity.Menu{}
	err := m.db.Table("menu").
		// Where("parent_id <> 0 AND status > 0").
		Where("parent_id IN ? AND status > 0", menuIDs).
		Find(&menus).Error
	return menus, err
}

func (m *MenuRepo) GetCategory(menuIDs []int) ([]entity.Category, error) {
	categorys := []entity.Category{}
	err := m.db.Table("product_category").
		// Where("status > 0 AND menu_id = ?", menuID).
		Where("status > 0 AND menu_id IN ?", menuIDs).
		Find(&categorys).Error
	return categorys, err
}

func (m *MenuRepo) GetProductLine(categoryID int) ([]entity.ProductLine, error) {
	productLine := []entity.ProductLine{}
	err := m.db.Table("product").
		Where("status > 0 AND category_id = ?", categoryID).
		Find(&productLine).Error
	return productLine, err
}
