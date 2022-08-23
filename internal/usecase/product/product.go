package usecase

import (
	"github.com/duykb2015/ecom-api/internal/entity"
)

type ProductUsecase struct {
	repo ProductRepo
}

func New(r ProductRepo) *ProductUsecase {
	return &ProductUsecase{r}
}

func (uc *ProductUsecase) GetAll() ([]entity.ProductLineRespond, error) {
	entities, err := uc.repo.GetAllProduct()
	if err != nil {
		return nil, err
	}
	product := []entity.ProductLineRespond{}

	for _, val := range entities {
		p := entity.ProductLineRespond{
			ID:     int(val.ID),
			Name:   val.Name,
			Slug:   val.Slug,
			Status: int(val.Status),
		}
		product = append(product, p)
	}

	return product, nil

}

func (uc *ProductUsecase) GetByCategory(id int) ([]entity.ProductLineRespond, error) {
	entities, err := uc.repo.GetAllProductByCategory(id)
	if err != nil {
		return nil, err
	}
	product := []entity.ProductLineRespond{}

	for _, val := range entities {
		p := entity.ProductLineRespond{
			ID:     int(val.ID),
			Name:   val.Name,
			Slug:   val.Slug,
			Status: int(val.Status),
		}
		product = append(product, p)
	}

	return product, nil
}

func (uc *ProductUsecase) Items(id int) ([]entity.ProductRespond, error) {
	entities, err := uc.repo.GetAllProductItemsByProduct(id)

	if err != nil {
		return nil, err
	}
	product := []entity.ProductRespond{}

	for _, val := range entities {
		p := entity.ProductRespond{
			ID:     val.ID,
			Name:   val.Name,
			Slug:   val.Slug,
			Status: val.Status,
			Items:  []entity.ProductItemsRespond{},
		}
		for _, item := range val.ProductItems {
			p.Items = append(p.Items, entity.ProductItemsRespond{
				ID:     item.ID,
				Name:   item.Name,
				Slug:   item.Slug,
				Status: item.Status,
			})
		}
		product = append(product, p)
	}

	return product, nil
}

func (uc *ProductUsecase) ItemInfo(product_id int, product_item_id int) (entity.ProductRespond, error) {
	entities, err := uc.repo.GetProductItemInfo(product_id, product_item_id)

	if err != nil {
		return entity.ProductRespond{}, err
	}
	product := entity.ProductRespond{
		ID:                   entities.ID,
		Name:                 entities.Name,
		Slug:                 entities.Slug,
		AdditionalInfomation: entities.AdditionalInfomation,
		SupportInfomation:    entities.SupportInfomation,
		Description:          entities.Description,
		Status:               entities.Status,
		Attributes:           []entity.ProductAttributesRespond{},
		Items:                []entity.ProductItemsRespond{},
	}

	return product, nil
}

//a
