package v1

import (
	"github.com/duykb2015/ecom-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, p usecase.ProductRepo, m usecase.MenuRepo) {

	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("/v1")
	{
		NewProductRoutes(h, p)
		NewMenuRoutes(h, m)
	}
}
