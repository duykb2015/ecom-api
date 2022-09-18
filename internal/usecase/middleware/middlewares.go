package middleware

import (
	"github.com/duykb2015/ecom-api/config"
)

type MiddleWareManager struct {
	cfg *config.App
}

func NewMiddleWareManager(cfg *config.App) *MiddleWareManager {
	return &MiddleWareManager{cfg: cfg}
}
