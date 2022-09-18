package main

import (
	"log"

	"github.com/duykb2015/ecom-api/config"
	"github.com/duykb2015/ecom-api/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("app - main - Load env error: %s", err)
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(cfg)
}
