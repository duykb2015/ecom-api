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
		Where("parent_id <> 0 AND id IN ? AND status > 0", menuIDs).
		Find(&menus).Error
	return menus, err
}

func (m *MenuRepo) GetCategory() ([]entity.Category, error) {
	return nil, nil
}

func (m *MenuRepo) GetProductLine() ([]entity.ProductLine, error) {
	return nil, nil
}
