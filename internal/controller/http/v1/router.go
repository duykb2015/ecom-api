package v1

import (
	menu "github.com/duykb2015/ecom-api/internal/usecase/menu"
	product "github.com/duykb2015/ecom-api/internal/usecase/product"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, p product.Product, m menu.Menu) {

	h := handler.Group("/v1")
	{
		NewProductRoutes(h, p)
		NewMenuRoutes(h, m)
	}
}
