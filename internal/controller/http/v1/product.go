package v1

import (
	"net/http"
	"time"

	"github.com/duykb2015/ecom-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	p usecase.ProductRepo
}

func newProductRoutes(handler *gin.RouterGroup, p usecase.ProductRepo) {
	r := &ProductRoutes{p: p}
	h := handler.Group("product")
	{
		h.GET("/get-all", r.getAllProduct)
	}
}

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

func (r *ProductRoutes) getAllProduct(c *gin.Context) {
	product, err := r.p.GetAllProduct()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)

}
