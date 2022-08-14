package entity

import "time"

type Product struct {
	ID                   uint      `gorm:"primary_key"`
	AdminID              uint      `gorm:"column:admin_id"`
	CategoryID           uint      `gorm:"column:category_id"`
	Name                 string    `gorm:"column:name"`
	Slug                 string    `gorm:"column:slug"`
	AdditionalInfomation string    `gorm:"column:additional_infomation"`
	SupportInfomation    string    `gorm:"column:support_infomation"`
	Description          string    `gorm:"column:description"`
	Status               uint      `gorm:"column:status"`
	CreatedAt            time.Time `gorm:"column:created_at"`
	UpdatedAt            time.Time `gorm:"column:updated_at"`
}
