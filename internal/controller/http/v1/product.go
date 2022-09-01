package v1

import (
	"net/http"
	"strconv"

	usecase "github.com/duykb2015/ecom-api/internal/usecase/product"
	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	p usecase.Product
}

func NewProductRoutes(handler *gin.RouterGroup, p usecase.Product) {
	r := &ProductRoutes{p: p}
	h := handler.Group("product")
	{
		h.GET("/", r.GetAllProduct)
		h.GET("/category/:id", r.getByCategory)
		h.GET("/item/:productID/:itemID", r.getItems)
	}
}

func (r *ProductRoutes) GetAllProduct(c *gin.Context) {
	product, err := r.p.Get()

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

func (r *ProductRoutes) getByCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	product, err := r.p.Category(categoryID)

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

func (r *ProductRoutes) getItems(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	itemID, err := strconv.Atoi(c.Param("itemID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	product, err := r.p.Items(productID, itemID)

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
