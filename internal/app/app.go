package app

import (
	"log"

	"github.com/duykb2015/ecom-api/config"
	v1 "github.com/duykb2015/ecom-api/internal/controller/http/v1"
	"github.com/duykb2015/ecom-api/internal/usecase"
	menu "github.com/duykb2015/ecom-api/internal/usecase/menu/repomysql"
	product "github.com/duykb2015/ecom-api/internal/usecase/product/repomysql"
	"github.com/duykb2015/ecom-api/pkg/db/mysql"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {

	gin.SetMode(gin.ReleaseMode)

	db, err := mysql.New(&cfg.MySQL)
	if err != nil {
		log.Fatalf("app - Run - mysql.New: %s", err)
	}

	productUsecase := usecase.NewProduct(product.NewProductRepo(db))
	menuUsecase := usecase.NewMenu(menu.NewMenuRepo(db))

	handler := gin.Default()
	v1.NewRouter(handler, productUsecase, menuUsecase)
	handler.Run(":8081")

}
