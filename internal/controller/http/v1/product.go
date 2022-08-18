package v1

import (
	"net/http"
	"strconv"

	usecase "github.com/duykb2015/ecom-api/internal/usecase/product"
	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	p usecase.ProductRepo
}
type Product struct {
	ID                   uint   `gorm:"primary_key"`
	CategoryID           uint   `gorm:"column:category_id"`
	Name                 string `gorm:"column:name"`
	Slug                 string `gorm:"column:slug"`
	AdditionalInfomation string `gorm:"column:additional_information"`
	SupportInfomation    string `gorm:"column:support_information"`
	Description          string `gorm:"column:description"`
	Status               uint   `gorm:"column:status"`
	ProductAttributes    []ProductAttributes
	ProductItems         []ProductItems
}

type ProductItems struct {
	ID                uint   `gorm:"primary_key"`
	ProductID         uint   `gorm:"column:product_id"`
	AdminID           uint   `gorm:"column:admin_id"`
	Name              string `gorm:"column:name"`
	Slug              string `gorm:"column:slug"`
	Status            uint   `gorm:"column:status"`
	ProductAttributes []ProductAttributes
	ProductItemColors []ProductItemColors
	ProductItemImages []ProductItemImages
}

type ProductAttributes struct {
	ID             uint   `gorm:"primary_key"`
	ProductID      uint   `gorm:"column:product_id"`
	ProductItemsID uint   `gorm:"column:product_item_id"`
	Name           string `gorm:"column:name"`
	Key            string `gorm:"column:key"`
	Value          string `gorm:"column:value"`
}

type ProductItemColors struct {
	ID             uint    `gorm:"primary_key"`
	ProductItemsID uint    `gorm:"column:product_item_id"`
	Name           string  `gorm:"column:name"`
	Hexcode        string  `gorm:"column:hexcode"`
	Price          float64 `gorm:"column:price"`
	Discount       float64 `gorm:"column:discount"`
	Quantity       uint    `gorm:"column:quantity"`
	Status         uint    `gorm:"column:status"`
}

type ProductItemImages struct {
	ID             uint   `gorm:"primary_key"`
	ProductItemsID uint   `gorm:"column:product_item_id"`
	Name           string `gorm:"column:name"`
}

func NewProductRoutes(handler *gin.RouterGroup, p usecase.ProductRepo) {
	r := &ProductRoutes{p: p}
	h := handler.Group("product")
	{
		h.GET("/", r.getAllProductLine)
		h.GET("/category/:id", r.getAllProductLineByCategory)
		h.GET("/line/:id", r.getAllProductItemsByProductLine)
		h.GET("/info/:product_id/:product_item_id", r.getProductItemInfo)
	}
}

func (r *ProductRoutes) getAllProductLine(c *gin.Context) {
	product, err := r.p.GetAllProductLine()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"result": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":  "success",
		"result": product,
	})

}

func (r *ProductRoutes) getAllProductLineByCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	product, err := r.p.GetAllProductLineByCategory(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":  "success",
		"result": product,
	})
}

func (r *ProductRoutes) getAllProductItemsByProductLine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	product, err := r.p.GetAllProductItemsByProductLine(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":  "success",
		"result": product,
	})
}

func (r *ProductRoutes) getProductItemInfo(c *gin.Context) {
	product_id, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"result": "",
		})
		return
	}

	product_item_id, err := strconv.Atoi(c.Param("product_item_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"result": "",
		})
		return
	}

	product, err := r.p.GetProductItemInfo(product_id, product_item_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"result": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":  "success",
		"result": product,
	})
}
