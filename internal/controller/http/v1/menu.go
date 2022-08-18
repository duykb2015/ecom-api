package v1

import (
	usecase "github.com/duykb2015/ecom-api/internal/usecase/menu"
	"github.com/gin-gonic/gin"
)

type MenuRoutes struct {
	m usecase.MenuRepo
}

func NewMenuRoutes(handler *gin.RouterGroup, m usecase.MenuRepo) {
	r := &MenuRoutes{m}
	h := handler.Group("/menu")
	{
		h.GET("/get-all", r.GetAllMenu)
	}
}

func (r *MenuRoutes) GetAllMenu(c *gin.Context) {
	menu, err := r.m.GetAllMenu()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, menu)
}
