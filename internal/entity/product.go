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
	ID                   uint
	Name                 string
	Slug                 string
	AdditionalInfomation string
	SupportInfomation    string
	Description          string
	Status               uint
	Attributes           []ProductAttributesResponse
	Items                []ProductItemsResponse
}
type ProductItemsResponse struct {
	ID         uint
	Name       string
	Slug       string
	Status     uint
	Attributes []ProductAttributesResponse
	Colors     []ProductItemColorsResponse
	Images     []ProductItemImagesResponse
}

type ProductAttributesResponse struct {
	ID     uint
	Name   string
	Key    string
	Value  string
	Status uint
}

type ProductItemColorsResponse struct {
	ID       uint
	Name     string
	Hexcode  string
	Price    float64
	Discount float64
	Quantity uint
	Status   uint
}

type ProductItemImagesResponse struct {
	ID   uint
	Name string
}
