package main

import (
	"log"

	"github.com/duykb2015/ecom-api/config"
	"github.com/duykb2015/ecom-api/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(cfg)
}
