package app

import (
	"github.com/duykb2015/ecom-api/config"
	v1 "github.com/duykb2015/ecom-api/internal/controller/http/v1"
	"github.com/duykb2015/ecom-api/internal/usecase/repomysql"
	"github.com/duykb2015/ecom-api/pkg/db/mysql"
	"github.com/gin-gonic/gin"
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
