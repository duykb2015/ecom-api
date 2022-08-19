package entity

import "time"

type Menu struct {
	ID        int       `gorm:"primary_key"`
	ParentID  int       `gorm:"column:parent_id"`
	Name      string    `gorm:"column:name"`
	Slug      string    `gorm:"column:slug"`
	Type      int       `gorm:"column:type"`
	Status    int       `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	SubMenu   []SubMenu
	Category  []Category
}
type SubMenu struct {
	ID        int       `gorm:"primary_key; column:sub_id"`
	MenuID    int       `gorm:"column:id"`
	ParentID  int       `gorm:"column:sub_parent_id"`
	Name      string    `gorm:"column:sub_name"`
	Slug      string    `gorm:"column:sub_slug"`
	Type      int       `gorm:"column:sub_type"`
	Status    int       `gorm:"column:sub_status"`
	CreatedAt time.Time `gorm:"column:sub_created_at"`
	UpdatedAt time.Time `gorm:"column:sub_updated_at"`
}
type Category struct {
	ID          int       `gorm:"primary_key; column:id"`
	MenuID      int       `gorm:"column:menu_id"`
	Name        string    `gorm:"column:name"`
	Slug        string    `gorm:"column:slug"`
	Status      int       `gorm:"column:status"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	ProductLine []ProductLine
}
type ProductLine struct {
	ID         int    `gorm:"primary_key; column:id"`
	CategoryID int    `gorm:"column:category_id"`
	Name       string `gorm:"column:name"`
	Slug       string `gorm:"column:slug"`
	Status     int    `gorm:"column:status"`
}

type MenuRespond struct {
	ID       int
	Name     string
	Slug     string
	SubMenu  []SubMenuRespond
	Category []CategoryRespond
}
type SubMenuRespond struct {
	ID   int
	Name string
	Slug string
}
type CategoryRespond struct {
	ID          int
	Name        string
	Slug        string
	ProductLine []ProductLineRespond
}
type ProductLineRespond struct {
	ID   int
	Name string
	Slug string
}
