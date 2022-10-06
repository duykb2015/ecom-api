package v1

import (
	"net/http"

	usecase "github.com/duykb2015/ecom-api/internal/usecase/menu"
	"github.com/duykb2015/ecom-api/pkg/httpclient"
	"github.com/gin-gonic/gin"
)

type MenuRoutes struct {
	m usecase.Menu
}

func NewMenuRoutes(handler *gin.RouterGroup, m usecase.Menu) {
	r := &MenuRoutes{m}
	h := handler.Group("/menu")
	{
		h.GET("", r.GetAllMenu)
	}
}

func (r *MenuRoutes) GetAllMenu(c *gin.Context) {
	menu, err := r.m.Get()
	if err != nil {
		httpclient.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	httpclient.NewResponse(c, http.StatusOK, "ok", menu)
}
