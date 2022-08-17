package v1

import (
	"net/http"

	"github.com/duykb2015/ecom-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	p usecase.ProductRepo
}

func NewProductRoutes(handler *gin.RouterGroup, p usecase.ProductRepo) {
	r := &ProductRoutes{p: p}
	h := handler.Group("product")
	{
		h.GET("/get-all", r.getAllProduct)
		h.GET("/get-by-slug/:slug", r.getProductBySlug)
		h.GET("/get-by-category/:slug", r.GetAllProductByCategory)
		h.GET("/get-by-product-line/:slug", r.GetAllProductByProductLine)
	}
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

func (r *ProductRoutes) GetAllProductByCategory(c *gin.Context) {
	slug := c.Param("slug")
	product, err := r.p.GetAllProductByCategory(slug)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (r *ProductRoutes) GetAllProductByProductLine(c *gin.Context) {
	slug := c.Param("slug")
	product, err := r.p.GetAllProductByProductLine(slug)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (r *ProductRoutes) getProductBySlug(c *gin.Context) {
	slug := c.Param("slug")
	product, err := r.p.GetProductBySlug(slug)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}
