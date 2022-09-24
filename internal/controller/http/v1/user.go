package v1

import (
	"net/http"

	"github.com/duykb2015/ecom-api/internal/controller/http/middleware"
	"github.com/duykb2015/ecom-api/internal/entity"
	usecase "github.com/duykb2015/ecom-api/internal/usecase/user"
	"github.com/duykb2015/ecom-api/pkg/httpclient"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	u usecase.User
}

func NewUserRoutes(handler *gin.RouterGroup, u usecase.User, mw *middleware.MiddleWareManager) {
	r := &UserRoutes{u}
	h := handler.Group("user")
	{
		h.POST("auth/login", r.Login)
		h.POST("auth/register", r.Register)
		h.Use(mw.AuthJWTMiddleWare())
	}
}

func (r *UserRoutes) Login(c *gin.Context) {
	request := entity.AuthRequest{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, httpclient.NewResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}
	user, err := r.u.AuthLogin(request.Email, request.Password)

	if err != nil {
		c.JSON(http.StatusOK, httpclient.NewResponse(http.StatusOK, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, httpclient.NewResponse(http.StatusOK, "ok", user))
}

func (r *UserRoutes) Register(c *gin.Context) {
	request := entity.AuthRequest{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, httpclient.NewResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}
	user, err := r.u.AuthRegister(request)

	if err != nil {
		c.JSON(http.StatusOK, httpclient.NewResponse(http.StatusOK, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, httpclient.NewResponse(http.StatusOK, "ok", user))
}
