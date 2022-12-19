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
	// categorys := []entity.Category{}
	if len(menuIDs) > 0 {
		childens, err = uc.repo.GetChildens(menuIDs)
		if err != nil {
			return resp, err
		}

		// categorys, err = uc.repo.GetCategory(menuIDs)
		// if err != nil {
		// 	return resp, err
		// }
	}

	for _, parent := range parents {
		subMenus := make([]entity.MenuResponse, 0)
		// categoryResp := make([]entity.CategoryResponse, 0)
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

		// for _, category := range categorys {
		// }

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

func (uc *MenuUsecase) _Get() ([]entity.MenuResponse, error) {
	resp := make([]entity.MenuResponse, 0)

	// Get parent
	parents, err := uc.repo.GetParents()
	if err != nil {
		return resp, nil
	}

	// Get childent
	childens, err := uc.repo.GetChildens(nil)
	if err != nil {
		return resp, err
	}

	for _, parent := range parents {
		subMenus := make([]entity.MenuResponse, 0)
		for _, childen := range childens {
			if childen.ParentID == parent.ID {
				category, err := uc.GetCategory(parent.ID)
				if err != nil {
					return resp, err
				}
				subMenus = append(subMenus, entity.MenuResponse{
					ID:        childen.ID,
					ParentID:  childen.ParentID,
					Name:      childen.Name,
					Slug:      childen.Slug,
					Type:      childen.Type,
					Categorys: category,
				})
			}
		}

		category, err := uc.GetCategory(parent.ID)
		if err != nil {
			return resp, err
		}

		resp = append(resp, entity.MenuResponse{
			ID:        parent.ID,
			ParentID:  parent.ParentID,
			Name:      parent.Name,
			Slug:      parent.Slug,
			Type:      parent.Type,
			SubMenus:  subMenus,
			Categorys: category,
		})
	}

	return resp, nil
}

func (uc *MenuUsecase) GetCategory(menuIDs int) ([]entity.CategoryResponse, error) {
	resp := make([]entity.CategoryResponse, 0)

	categorys, err := uc.repo.GetCategory(nil)
	if err != nil {
		return resp, err
	}

	for _, category := range categorys {
		resp = append(resp, entity.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
			Slug: category.Slug,
		})
	}

	return resp, nil
}
