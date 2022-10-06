package entity

import "time"

type Cart struct {
	ID            int       `gorm:"primary_key"`
	UserID        int       `gorm:"user_id"`
	ProductItemID int       `gorm:"product_item_id"`
	ColorID       int       `gorm:"color"`
	Quantity      int       `gorm:"quantity"`
	Status        int       `gorm:"column:status"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

type CartResponse struct {
	ID            int                  `json:"id"`
	UserID        int                  `json:"user_id"`
	ProductItem   ProductItemsResponse `json:"product_item"`
	SelectedColor int                  `json:"selected_color"`
	Status        int                  `json:"status"`
}
