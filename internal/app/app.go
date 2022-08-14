package app

import (
	v1 "github.com/duykb2015/login-api/internal/controller/http/v1"
	"github.com/duykb2015/login-api/internal/usecase/repo"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)

	dns := "root:1@tcp(localhost:3306)/ecommerce?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	handler := gin.Default()
	v1.NewRouter(handler, repo.New(db))
	handler.Run(":8080")
}
