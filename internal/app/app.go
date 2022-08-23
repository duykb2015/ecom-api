package app

import (
	"log"

	"github.com/duykb2015/ecom-api/config"
	v1 "github.com/duykb2015/ecom-api/internal/controller/http/v1"
	menuUsecase "github.com/duykb2015/ecom-api/internal/usecase/menu"
	menuRepo "github.com/duykb2015/ecom-api/internal/usecase/menu/repomysql"
	productUsecase "github.com/duykb2015/ecom-api/internal/usecase/product"
	productRepo "github.com/duykb2015/ecom-api/internal/usecase/product/repomysql"
	"github.com/duykb2015/ecom-api/pkg/db/mysql"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {

	gin.SetMode(gin.ReleaseMode)

	db, err := mysql.New(&cfg.MySQL)
	if err != nil {

		log.Fatalf("app - Run - mysql.New: %s", err)
	}

	productUC := productUsecase.NewProduct(productRepo.NewProductRepo(db))
	menuUC := menuUsecase.New(menuRepo.New(db))

	handler := gin.Default()
	v1.NewRouter(handler, productUC, menuUC)
	handler.Run(":8080")
}
