package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

// Xem lại cách đặt tên
type (
	MenuRepo interface {
		GetParents() ([]entity.Menu, error)
		GetChildens(menuIDs []int) ([]entity.Menu, error)
	}
	Menu interface {
		Get() ([]entity.MenuResponse, error)
	}
)
