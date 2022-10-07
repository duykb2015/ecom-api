package v1

import (
	"net/http"
	"strconv"

	usecase "github.com/duykb2015/ecom-api/internal/usecase/product"
	"github.com/duykb2015/ecom-api/pkg/httpclient"
	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	p usecase.Product
}

func NewProductRoutes(handler *gin.RouterGroup, p usecase.Product) {
	r := &ProductRoutes{p: p}
	h := handler.Group("product")
	{
		h.GET("", r.GetAllProduct)
		h.GET("/category/:id", r.getByCategory)
		h.GET("/hotdeal", r.GetHotDeal)
		h.GET("/line", r.GetLine)
	}
}

func (r *ProductRoutes) GetAllProduct(c *gin.Context) {
	product, err := r.p.Get()
	if err != nil {
		httpclient.NewResponse(c, http.StatusBadRequest, "error", err)
		return
	}
	httpclient.NewResponse(c, http.StatusOK, "ok", product)

}

func (r *ProductRoutes) getByCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httpclient.NewResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	product, err := r.p.Category(categoryID)
	if err != nil {
		httpclient.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	httpclient.NewResponse(c, http.StatusOK, "ok", product)
}

func (r *ProductRoutes) GetHotDeal(c *gin.Context) {
	product, err := r.p.GetHotDeal()
	if err != nil {
		httpclient.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	httpclient.NewResponse(c, http.StatusOK, "ok", product)
}

func (r *ProductRoutes) GetLine(c *gin.Context) {
	product, err := r.p.GetLine()
	if err != nil {
		httpclient.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	httpclient.NewResponse(c, http.StatusOK, "ok", product)
}
