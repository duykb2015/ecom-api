package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type (
	// ProductUsecase interface
	MenuRepo interface {
		GetAllMenu() ([]entity.Menu, error)
	}
)
