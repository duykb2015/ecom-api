package entity

import "time"

type Product struct {
	ID                   uint      `gorm:"primary_key"`
	AdminID              uint      `gorm:"column:admin_id"`
	CategoryID           uint      `gorm:"column:category_id"`
	Name                 string    `gorm:"column:name"`
	Slug                 string    `gorm:"column:slug"`
	AdditionalInfomation string    `gorm:"column:additional_information"`
	SupportInfomation    string    `gorm:"column:support_information"`
	Description          string    `gorm:"column:description"`
	Status               uint      `gorm:"column:status"`
	CreatedAt            time.Time `gorm:"column:created_at"`
	UpdatedAt            time.Time `gorm:"column:updated_at"`
	ProductAttributes    []ProductAttributes
	ProductItems         []ProductItems
}

type ProductItems struct {
	ID                uint      `gorm:"primary_key"`
	ProductID         uint      `gorm:"column:product_id"`
	AdminID           uint      `gorm:"column:admin_id"`
	Name              string    `gorm:"column:name"`
	Slug              string    `gorm:"column:slug"`
	Status            uint      `gorm:"column:status"`
	CreatedAt         time.Time `gorm:"column:created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at"`
	ProductAttributes []ProductAttributes
	ProductItemColors []ProductItemColors
	ProductItemImages []ProductItemImages
}

type ProductAttributes struct {
	ID             uint      `gorm:"primary_key"`
	ProductID      uint      `gorm:"column:product_id"`
	ProductItemsID uint      `gorm:"column:product_item_id"`
	Name           string    `gorm:"column:name"`
	Key            string    `gorm:"column:key"`
	Value          string    `gorm:"column:value"`
	Status         uint      `gorm:"column:status"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

type ProductItemColors struct {
	ID             uint      `gorm:"primary_key"`
	ProductItemsID uint      `gorm:"column:product_item_id"`
	Name           string    `gorm:"column:name"`
	Hexcode        string    `gorm:"column:hexcode"`
	Price          float64   `gorm:"column:price"`
	Discount       float64   `gorm:"column:discount"`
	Quantity       uint      `gorm:"column:quantity"`
	Status         uint      `gorm:"column:status"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

type ProductItemImages struct {
	ID             uint      `gorm:"primary_key"`
	ProductItemsID uint      `gorm:"column:product_item_id"`
	Name           string    `gorm:"column:name"`
	Status         uint      `gorm:"column:status"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

type ProductResponse struct {
	ID                    uint                        `json:"id"`
	Name                  string                      `json:"name"`
	Slug                  string                      `json:"slug"`
	AdditionalInformation string                      `json:"addtional_information"`
	SupportInformation    string                      `json:"support_information"`
	Description           string                      `json:"description"`
	Status                uint                        `json:"status"`
	Attributes            []ProductAttributesResponse `json:"attributes"`
	Items                 []ProductItemsResponse      `json:"items"`
	BasicItemInfo         []ProductBasicInfoResponse  `json:"items_info"`
}
type ProductItemsResponse struct {
	ID         uint                        `json:"id"`
	Name       string                      `json:"name"`
	Slug       string                      `json:"slug"`
	Status     uint                        `json:"status"`
	Attributes []ProductAttributesResponse `json:"attributes"`
	Colors     []ProductItemColorsResponse `json:"colors"`
	Images     []ProductItemImagesResponse `json:"images"`
}

type ProductAttributesResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Key    string `json:"key"`
	Value  string `json:"value"`
	Status uint   `json:"status"`
}

type ProductItemColorsResponse struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Hexcode  string  `json:"hexcode"`
	Price    float64 `json:"price"`
	Discount float64 `json:"discount"`
	Quantity uint    `json:"quantity"`
	Status   uint    `json:"status"`
}

type ProductItemImagesResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ProductBasicInfoResponse struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	Image         string `json:"image"`
	Price         uint   `json:"price"`
	Discount      uint   `json:"discount"`
	DiscountPrice uint   `json:"discount_price"`
}
