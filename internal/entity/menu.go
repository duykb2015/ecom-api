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
	ID     int
	Name   string
	Slug   string
	Status int
}

type MenuResponse struct {
	ID       int            `json:"id"`
	ParentID int            `json:"parent_id"`
	Name     string         `json:"name"`
	Slug     string         `json:"slug"`
	Type     int            `json:"type"`
	SubMenus []MenuResponse `json:"sub_menu"`
}
