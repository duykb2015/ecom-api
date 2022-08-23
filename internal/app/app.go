package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/duykb2015/ecom-api/config"
	v1 "github.com/duykb2015/ecom-api/internal/controller/http/v1"
	menuUsecase "github.com/duykb2015/ecom-api/internal/usecase/menu"
	menuRepo "github.com/duykb2015/ecom-api/internal/usecase/menu/repomysql"
	productUsecase "github.com/duykb2015/ecom-api/internal/usecase/product"
	productRepo "github.com/duykb2015/ecom-api/internal/usecase/product/repomysql"
	"github.com/duykb2015/ecom-api/pkg/httpserver"
	"github.com/duykb2015/ecom-api/pkg/logger"
	"github.com/duykb2015/ecom-api/pkg/mysql"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {

	gin.SetMode(gin.ReleaseMode)
	l := logger.New(cfg.Log.Level)

	db, err := mysql.New(&cfg.MySQL)
	if err != nil {
		log.Fatalf("app - Run - mysql.New: %s", err)
	}

	productUC := productUsecase.NewProduct(productRepo.NewProductRepo(db))
	menuUC := menuUsecase.NewMenu(menuRepo.NewMenuRepo(db))

	handler := gin.Default()
	v1.NewRouter(handler, productUC, menuUC)
	handler.Run(":8080")

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
