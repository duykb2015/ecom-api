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
	// SubMenu   []SubMenu `gorm:"foreignkey:SubID; references: ID"`
}

type SubMenu struct {
	SubID        int       `gorm:"primary_key; column:sub_id"`
	SubParentID  int       `gorm:"column:sub_parent_id"`
	SubName      string    `gorm:"column:sub_name"`
	SubSlug      string    `gorm:"column:sub_slug"`
	SubType      int       `gorm:"column:sub_type"`
	SubStatus    int       `gorm:"column:sub_status"`
	SubCreatedAt time.Time `gorm:"column:sub_created_at"`
	SubUpdatedAt time.Time `gorm:"column:sub_updated_at"`
}

type Category struct {
	CategoryID        int       `gorm:"primary_key; column:category_id"`
	CategoryMenuID    int       `gorm:"column:category_menu_id"`
	CategoryName      string    `gorm:"column:category_name"`
	CategorySlug      string    `gorm:"column:category_slug"`
	CategoryStatus    int       `gorm:"column:category_status"`
	CategoryCreatedAt time.Time `gorm:"column:category_created_at"`
	CategoryUpdatedAt time.Time `gorm:"column:category_updated_at"`
}
