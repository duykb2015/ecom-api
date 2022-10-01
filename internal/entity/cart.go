package entity

import "time"

type Cart struct {
	ID            uint      `gorm:"primary_key"`
	UserID        string    `gorm:"user_id"`
	ProductItemId string    `gorm:"product_item_id"`
	ProductName   string    `gorm:"product_name"`
	ProductImage  string    `gorm:"product_image,omitempty"`
	Color         string    `gorm:"color"`
	Discount      int       `gorm:"discount"`
	Quantity      int       `gorm:"quantity"`
	Price         int       `gorm:"price"`
	Status        uint      `gorm:"column:status"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

type CartResponse struct {
	ID            uint   `json:"id"`
	UserID        string `json:"user_id"`
	ProductItemId string `json:"product_item_id"`
	ProductName   string `json:"product_name"`
	ProductImage  string `json:"product_image,omitempty"`
	Color         string `json:"color"`
	Discount      int    `json:"discount"`
	Quantity      int    `json:"quantity"`
	Price         int    `json:"price"`
	Status        uint   `json:"status"`
}
