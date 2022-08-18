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

func NewProductRoutes(handler *gin.RouterGroup, p usecase.ProductRepo) {
	r := &ProductRoutes{p: p}
	h := handler.Group("product")
	{
		h.GET("/get-all", r.getAllProductLine)
		h.GET("/get-by-category/:slug", r.getAllProductLineByCategory)
		h.GET("/get-by-line/:id", r.getAllProductItemsByProductLine)
		h.GET("/get-by-slug/:id/:slug", r.getProductItemBySlug)
	}
}

func (r *ProductRoutes) getAllProductLine(c *gin.Context) {
	product, err := r.p.GetAllProductLine()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)

}

func (r *ProductRoutes) getAllProductLineByCategory(c *gin.Context) {
	slug := c.Param("slug")
	product, err := r.p.GetAllProductLineByCategory(slug)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (r *ProductRoutes) getAllProductItemsByProductLine(c *gin.Context) {
	slug, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	product, err := r.p.GetAllProductItemsByProductLine(slug)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (r *ProductRoutes) getProductItemBySlug(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	slug := c.Param("slug")
	product, err := r.p.GetProductItemBySlug(id, slug)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}
