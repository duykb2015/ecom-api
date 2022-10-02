package httpclient

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	MessageOK = "Success"
)

const (
	ErrBadRequest       = "Bad request"
	ErrNoSuchUser       = "user_id not found"
	ErrNotFound         = "Not Found"
	ErrUnauthorized     = "Unauthorized"
	ErrForbidden        = "Forbidden"
	ErrBadQueryParams   = "Invalid query params"
	ErrPasswordNotMatch = "Password does not match"
)

var (
	Unauthorized     = errors.New("unauthorized")
	Invalid          = errors.New("invalid")
	InvalidJWTToken  = errors.New("invalid jwt token")
	InvalidJWTClaims = errors.New("invalid jwt claims")
)

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

func NewResponseWithGin(c *gin.Context, code int, msg string, result interface{}) {
	c.JSON(code, Response{
		Code:    code,
		Message: msg,
		Result:  result,
	})
}
