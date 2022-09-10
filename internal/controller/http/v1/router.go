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

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func NewResponse(code int, msg string, result interface{}) Response {
	return Response{
		Code:    code,
		Message: msg,
		Result:  result,
	}
}
