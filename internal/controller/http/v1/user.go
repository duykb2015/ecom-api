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
		h.POST("login", r.Login)
		h.POST("register", r.Register)
		h.Use(mw.AuthJWTMiddleWare())
		h.GET("getinfo", r.GetInfo)
		h.PUT("update", r.UpdateInfo)
		h.GET("cart", r.GetCart)
		h.POST("cart/add", r.AddItemToCart)
		h.PUT("cart/update", r.UpdateCart)
		h.DELETE("cart/delete", r.DeleteCart)
	}
}

func (r *UserRoutes) Login(c *gin.Context) {
	request := entity.AuthRequest{}
	if err := c.ShouldBind(&request); err != nil {
		httpclient.NewResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	user, err := r.u.AuthLogin(request.Email, request.Password)
	if err != nil {
		httpclient.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	httpclient.NewResponse(c, http.StatusOK, "ok", user)
}

func (r *UserRoutes) Register(c *gin.Context) {
	request := entity.AuthRequest{}
	if err := c.ShouldBind(&request); err != nil {
		httpclient.NewResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	user, err := r.u.AuthRegister(request)
	if err != nil {
		httpclient.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	httpclient.NewResponse(c, http.StatusOK, "ok", user)
}

func (r *UserRoutes) GetInfo(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		httpclient.NewResponse(c, http.StatusBadRequest, "No token header found", nil)
		return
	}

	user, err := r.u.GetInfo(tokenString)
	if err != nil {
		httpclient.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	httpclient.NewResponse(c, http.StatusOK, "ok", user)
}

func (r *UserRoutes) UpdateInfo(c *gin.Context) {
	request := entity.AuthRequest{}
	if err := c.ShouldBind(&request); err != nil {
		httpclient.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err := r.u.UpdateInfo(request)
	if err != nil {
		httpclient.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	httpclient.NewResponse(c, http.StatusOK, "ok", nil)
}

func (r *UserRoutes) GetCart(c *gin.Context) {
	userID := int(c.MustGet("user_id").(float64))
	cart, err := r.u.GetCart(userID)
	if err != nil {
		httpclient.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	httpclient.NewResponse(c, http.StatusOK, "ok", cart)
}

func (r *UserRoutes) AddItemToCart(c *gin.Context) {
	userID := int(c.MustGet("user_id").(float64))
	cart, err := r.u.GetCart(userID)
	if err != nil {
		httpclient.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	httpclient.NewResponse(c, http.StatusOK, "ok", cart)
}

func (r *UserRoutes) UpdateCart(c *gin.Context) {
	userID := int(c.MustGet("user_id").(float64))
	cart, err := r.u.GetCart(userID)
	if err != nil {
		httpclient.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	httpclient.NewResponse(c, http.StatusOK, "ok", cart)
}

func (r *UserRoutes) DeleteCart(c *gin.Context) {
	userID := int(c.MustGet("user_id").(float64))
	cart, err := r.u.GetCart(userID)
	if err != nil {
		httpclient.NewResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	httpclient.NewResponse(c, http.StatusOK, "ok", cart)
}
