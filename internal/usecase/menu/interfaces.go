package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type (
	MenuRepo interface {
		GetAllMenu() ([]entity.Menu, error)
	}
	Menu interface {
		MenuRespond() ([]entity.MenuRespond, error)
	}
)
