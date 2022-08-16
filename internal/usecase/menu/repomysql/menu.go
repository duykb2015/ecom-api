package repomysql

import (
	"github.com/duykb2015/ecom-api/internal/entity"
	"gorm.io/gorm"
)

type MenuRepo struct {
	db *gorm.DB
}

func NewMenuRepo(db *gorm.DB) *MenuRepo {
	return &MenuRepo{db}
}

func (m *MenuRepo) GetAllMenu() ([]entity.Menu, error) {
	menu := []entity.Menu{}
	m.db.Table("menu").Where("parent_id", 0).Find(&menu)
	for i, val := range menu {
		m.db.Table("product_category").Where("menu_id", val.ID).Find(&menu[i].Category)
		for j, category := range menu[i].Category {
			m.db.Table("product").Select("id, name, slug, status").Where("category_id", category.ID).Find(&menu[i].Category[j].ProductLine)
		}
	}
	return menu, nil
}
