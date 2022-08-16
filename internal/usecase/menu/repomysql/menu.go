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
	entities := []entity.Menu{}
	m.db.Table("menu").
		Joins("INNER JOIN menu submenu ON menu.id = submenu.parent_id").
		Joins("INNER JOIN product_category pc ON pc.menu_id = menu.id").
		Select("menu.id as id, menu.parent_id as parent_id, menu.name as name, menu.slug as slug, menu.type as type, menu.status as status, menu.created_at as created_at, menu.updated_at as updated_at, submenu.id as sub_id, submenu.parent_id as sub_parent_id, submenu.name as sub_name, submenu.slug as sub_slug, submenu.type as sub_type, submenu.status as sub_status, submenu.created_at as sub_created_at, submenu.updated_at as sub_updated_at, pc.id as category_id, pc.menu_id as category_menu_id, pc.name as category_name, pc.slug as category_slug, pc.status as category_status, pc.created_at as category_created_at, pc.updated_at as category_updated_at").
		Find(&entities)
	// m.db.Table("menu").Find(&entities)
	return entities, nil
}

// menusParent =  select * from menu where parent_id <> 0

// idsMenuParent = []string{}
// for _, val  := range menusParent{
// 	idsMenuParent = append(idsMenuParent, val.id)
// }

// menusChilden = select * from menu where parent_id in (idsMenuParent)

// menus := []MenuStruct{}
// for _, parent  := range menusParent{
// 	menus = MenuStruct{
// 		id: parent.id,
// 		....
// 	}
// 	for _, childen  := range menusChilden{
// 		if parent.id == childen.id {
// 			menus.sub_menus = append (menus.sub_menus, MenuStruct{
// 				id: parent.id,
// 				....
// 			})
// 		}
// 	}
// }
