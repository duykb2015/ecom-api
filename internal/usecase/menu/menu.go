package usecase

import (
	"github.com/duykb2015/ecom-api/internal/entity"
)

type MenuUsecase struct {
	repo MenuRepo
}

func New(r MenuRepo) *MenuUsecase {
	return &MenuUsecase{r}
}

func (uc *MenuUsecase) Get() ([]entity.MenuResponse, error) {
	resp := make([]entity.MenuResponse, 0) // tạo một slice entity.MenuResponse rỗng

	// Get parent
	parents, err := uc.repo.GetParents()
	if err != nil {
		return resp, nil
	}

	// Get childen
	menuIDs := []int{}
	for _, parent := range parents {
		menuIDs = append(menuIDs, parent.ID)
	}

	childens := []entity.Menu{}
	if len(menuIDs) > 0 {
		childens, err = uc.repo.GetChildens(menuIDs)
		if err != nil {
			return resp, err
		}
	}

	for _, parent := range parents {
		subMenus := make([]entity.MenuResponse, 0)
		for _, childen := range childens {
			if childen.ParentID == parent.ID {
				subMenus = append(subMenus, entity.MenuResponse{
					ID:       childen.ID,
					ParentID: childen.ParentID,
					Name:     childen.Name,
					Slug:     childen.Slug,
					Type:     childen.Type,
				})
			}
		}

		resp = append(resp, entity.MenuResponse{
			ID:       parent.ID,
			ParentID: parent.ParentID,
			Name:     parent.Name,
			Slug:     parent.Slug,
			Type:     parent.Type,
			SubMenus: subMenus,
		})
	}

	return resp, nil
}

func (uc *MenuUsecase) GetCategory() ([]entity.CategoryResponse, error)
