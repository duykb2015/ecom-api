package v1

import (
	"github.com/duykb2015/ecom-api/internal/controller/http/middleware"
	menu "github.com/duykb2015/ecom-api/internal/usecase/menu"
	product "github.com/duykb2015/ecom-api/internal/usecase/product"
	user "github.com/duykb2015/ecom-api/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, p product.Product, m menu.Menu, mw *middleware.MiddleWareManager, u user.User) {

	h := handler.Group("/v1")
	{
		NewProductRoutes(h, p)
		NewMenuRoutes(h, m)
		NewUserRoutes(h, u, mw)
	}
}
