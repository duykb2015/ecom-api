package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type MenuUsecase struct {
	repo MenuRepo
}

func NewMenu(r MenuRepo) *MenuUsecase {
	return &MenuUsecase{r}
}

func (m *MenuUsecase) GetAllMenu() ([]entity.MenuRespond, error) {
	return ConvertData(m.repo.GetAllMenu())
}

func ConvertData(entites []entity.Menu, err error) ([]entity.MenuRespond, error) {
	menu := []entity.MenuRespond{}
	for _, val := range entites {
		m := entity.MenuRespond{
			ID:       val.ID,
			Name:     val.Name,
			Slug:     val.Slug,
			SubMenu:  []entity.SubMenuRespond{},
			Category: []entity.CategoryRespond{},
		}
		for _, submenu := range val.SubMenu {
			m.SubMenu = append(m.SubMenu, entity.SubMenuRespond{
				ID:   submenu.ID,
				Name: submenu.Name,
				Slug: submenu.Slug,
			})
		}
		for key, category := range val.Category {
			m.Category = append(m.Category, entity.CategoryRespond{
				ID:          category.ID,
				Name:        category.Name,
				Slug:        category.Slug,
				ProductLine: []entity.ProductLineRespond{},
			})
			for _, product := range category.ProductLine {
				m.Category[key].ProductLine = append(m.Category[key].ProductLine, entity.ProductLineRespond{
					ID:   product.ID,
					Name: product.Name,
					Slug: product.Slug,
				})
			}
		}
		menu = append(menu, m)
	}
	return menu, nil
}
