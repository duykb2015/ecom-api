package app

import (
	"github.com/duykb2015/login-api/config"
	v1 "github.com/duykb2015/login-api/internal/controller/http/v1"
	"github.com/duykb2015/login-api/internal/usecase/repomysql"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
)

func Run(cfg *config.Config) {
	gin.SetMode(gin.ReleaseMode)

	db, err := mysql.New(&cfg.MySQL)
	if err != nil {
		panic(err.Error())
	}

	handler := gin.Default()
	v1.NewRouter(handler, repomysql.New(db))
	handler.Run(":8080")
}
