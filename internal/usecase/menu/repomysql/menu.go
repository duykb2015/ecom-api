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

func (m *MenuRepo) _GetAllMenu() ([]entity.Menu, error) {
	menu := []entity.Menu{}
	m.db.Table("menu").Where("parent_id = ? AND status > ?", 0, 0).Find(&menu)
	for i, val := range menu {
		m.db.Table("menu submenu").
			Select("submenu.id as sub_id, submenu.parent_id as sub_parent_id, submenu.name as sub_name, submenu.slug as sub_slug, submenu.type as sub_type, submenu.status as sub_status, submenu.created_at as sub_created_at, submenu.updated_at as sub_updated_at").
			Where("submenu.parent_id = ? AND submenu.status > ?", val.ID, 0).
			Find(&menu[i].SubMenu)
		m.db.Table("product_category").Where("menu_id", val.ID).Find(&menu[i].Category)
		for j, category := range menu[i].Category {
			m.db.Table("product").Select("id, name, slug, status").Where("category_id", category.ID).Find(&menu[i].Category[j].ProductLine)
		}
	}
	return menu, nil
}
func (m *MenuRepo) GetAllMenu() ([]entity.MenuRespond, error) {
	menu := []entity.Menu{}
	m.db.Table("menu").Where("parent_id = ? AND status > ?", 0, 0).Find(&menu)
	for i, val := range menu {
		m.db.Table("menu submenu").
			Select("submenu.id as sub_id, submenu.parent_id as sub_parent_id, submenu.name as sub_name, submenu.slug as sub_slug, submenu.type as sub_type, submenu.status as sub_status, submenu.created_at as sub_created_at, submenu.updated_at as sub_updated_at").
			Where("submenu.parent_id = ? AND submenu.status > ?", val.ID, 0).
			Find(&menu[i].SubMenu)
		m.db.Table("product_category").Where("menu_id", val.ID).Find(&menu[i].Category)
		for j, category := range menu[i].Category {
			m.db.Table("product").Select("id, name, slug, status").Where("category_id", category.ID).Find(&menu[i].Category[j].ProductLine)
		}
	}
	result := ConvertData(menu)
	return result, nil
}
func ConvertData(data []entity.Menu) []entity.MenuRespond {
	menu := []entity.MenuRespond{}
	for _, val := range data {
		m := entity.MenuRespond{
			Id:       val.ID,
			Name:     val.Name,
			Slug:     val.Slug,
			SubMenu:  []entity.SubMenuRespond{},
			Category: []entity.CategoryRespond{},
		}
		for _, submenu := range val.SubMenu {
			m.SubMenu = append(m.SubMenu, entity.SubMenuRespond{
				Id:   submenu.ID,
				Name: submenu.Name,
				Slug: submenu.Slug,
			})
		}
		for key, category := range val.Category {
			m.Category = append(m.Category, entity.CategoryRespond{
				Id:          category.ID,
				Name:        category.Name,
				Slug:        category.Slug,
				ProductLine: []entity.ProductLineRespond{},
			})
			for _, product := range category.ProductLine {
				m.Category[key].ProductLine = append(m.Category[key].ProductLine, entity.ProductLineRespond{
					Id:   product.ID,
					Name: product.Name,
					Slug: product.Slug,
				})
			}
		}
		menu = append(menu, m)
	}
	return menu
}
