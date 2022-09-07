package v1

import (
	"net/http"

	usecase "github.com/duykb2015/ecom-api/internal/usecase/menu"
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
	menu, _ := r.m.Get()
	c.JSON(http.StatusOK, NewResponse(http.StatusOK, "ok", menu))
}
