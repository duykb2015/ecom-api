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
	h := handler.Group("/user")
	{
		h.GET("/auth/login", r.Login)
		h.GET("/auth/login", r.Register)
		h.Use(mw.AuthJWTMiddleWare())
	}
}

func (r *UserRoutes) Login(c *gin.Context) {
	user, err := r.u.AuthLogin()

	if err != nil {
		c.JSON(http.StatusOK, httpclient.NewResponse(http.StatusOK, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, httpclient.NewResponse(http.StatusOK, "ok", user))
}

func (r *UserRoutes) Register(c *gin.Context) {
	user, err := r.u.AuthRegister()

	if err != nil {
		c.JSON(http.StatusOK, httpclient.NewResponse(http.StatusOK, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, httpclient.NewResponse(http.StatusOK, "ok", user))
}
