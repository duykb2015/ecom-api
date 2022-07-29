package v1

import (
	"github.com/duykb2015/login-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, p usecase.ProductRepo) {

	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("/v1")
	{
		newProductRoutes(h, p)
	}
}
