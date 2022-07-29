package app

import (
	"database/sql"

	v1 "github.com/duykb2015/login-api/internal/controller/http/v1"
	"github.com/duykb2015/login-api/internal/usecase/repo"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)

	db, err := sql.Open("mysql", "root:@/ecommercial")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	handler := gin.Default()
	v1.NewRouter(handler, repo.New(db))
	handler.Run(":8080")
}
