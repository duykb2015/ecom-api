package middleware

import (
	"fmt"
	_ "log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/duykb2015/ecom-api/pkg/httpclient"
	"github.com/gin-gonic/gin"
)

func (mw *MiddleWareManager) AuthJWTMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if err := mw.validateJWT(tokenString, ctx); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, httpclient.NewResponse(http.StatusUnauthorized, err.Error(), nil))
		}
		ctx.Next()
	}
}

func (mw *MiddleWareManager) validateJWT(tokenString string, ctx *gin.Context) error {
	if tokenString == "" {
		return httpclient.InvalidJWTToken
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method %v", t.Header["alg"])
		}
		secretKey := []byte(mw.cfg.JWTSecretKey)
		return secretKey, nil
	})
	if err != nil {
		if err.Error() != "Token is expired" {
			return err
		}
		token.Valid = true
		err = nil
	}

	if !token.Valid {
		return httpclient.InvalidJWTToken
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["user_id"]
		ctx.Set("user_id", userID)
	}

	return nil
}
