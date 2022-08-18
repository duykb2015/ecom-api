package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type MenuUsecase struct {
	repo MenuRepo
}

func NewMenu(r MenuRepo) *MenuUsecase {
	return &MenuUsecase{r}
}

func (m *MenuUsecase) GetAllMenu() ([]entity.MenuRespond, error) {
	return m.repo.GetAllMenu()
}
