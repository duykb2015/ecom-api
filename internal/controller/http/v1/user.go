package v1

import (
	"net/http"

	"github.com/duykb2015/ecom-api/internal/controller/http/middleware"
	usecase "github.com/duykb2015/ecom-api/internal/usecase/user"
	"github.com/duykb2015/ecom-api/pkg/httpclient"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	u usecase.User
}

func NewUserRoutes(handler *gin.RouterGroup, u usecase.User, mw *middleware.MiddleWareManager) {
	r := &UserRoutes{u}
	h := handler.Group("/auth/login")
	{
		h.GET("", r.Login)
	}
}

func (r *UserRoutes) Login(c *gin.Context) {
	menu, err := r.u.AuthLogin()
	if err != nil {
		c.JSON(http.StatusOK, httpclient.NewResponse(http.StatusOK, "error", err))
		return
	}
	c.JSON(http.StatusOK, httpclient.NewResponse(http.StatusOK, "ok", menu))
}
