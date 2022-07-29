package v1

import (
	"net/http"

	"github.com/duykb2015/login-api/internal/usecase"
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
	ID                string `json:"id" binding:"required"`
	Name              string `json:"name" binding:"required"`
	Slug              string `json:"slug" binding:"required"`
	Descriptions      string `json:"descriptions" binding:"required"`
	ShortDescriptions string `json:"short_descriptions" binding:"required"`
	Images            string `json:"images" binding:"required"`
	Price             int    `json:"price" binding:"required"`
	Quantity          int    `json:"quantity" binding:"required"`
	Status            int    `json:"status" binding:"required"`
	CreatedAt         string `json:"created_at" binding:"required"`
	UpdatedAt         string `json:"updated_at" binding:"required"`
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
