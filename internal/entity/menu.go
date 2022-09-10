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

type MenuResponse struct {
	ID        int                `json:"id"`
	ParentID  int                `json:"parent_id"`
	Name      string             `json:"name"`
	Slug      string             `json:"slug"`
	Type      int                `json:"type"`
	SubMenus  []MenuResponse     `json:"sub_menu"`
	Categorys []CategoryResponse `json:"category"`
}

type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ProductLineResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Status int    `json:"status"`
}
