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
		h.GET("/", r.getAll)
		h.GET("/category/:id", r.getByCategory)
		h.GET("/line/:id", r.getAllItems)
		h.GET("/info/:product_id/:product_item_id", r.getItemInfo)
	}
}

func (r *ProductRoutes) getAll(c *gin.Context) {
	product, err := r.p.GetAll()

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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	product, err := r.p.GetByCategory(id)

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

func (r *ProductRoutes) getAllItems(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	product, err := r.p.Items(id)

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

func (r *ProductRoutes) getItemInfo(c *gin.Context) {
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

	product, err := r.p.ItemInfo(product_id, product_item_id)
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
