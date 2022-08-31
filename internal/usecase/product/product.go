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

func (uc *ProductUsecase) GetAll() ([]entity.ProductLineResponse, error) {
	entities, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	product := []entity.ProductLineResponse{}

	for _, val := range entities {
		p := entity.ProductLineResponse{
			ID:     int(val.ID),
			Name:   val.Name,
			Slug:   val.Slug,
			Status: int(val.Status),
		}
		product = append(product, p)
	}

	return product, nil

}

func (uc *ProductUsecase) ByCategory(id int) ([]entity.ProductLineResponse, error) {
	entities, err := uc.repo.GetByCategory(id)
	if err != nil {
		return nil, err
	}
	product := []entity.ProductLineResponse{}

	for _, val := range entities {
		p := entity.ProductLineResponse{
			ID:     int(val.ID),
			Name:   val.Name,
			Slug:   val.Slug,
			Status: int(val.Status),
		}
		product = append(product, p)
	}

	return product, nil
}

func (uc *ProductUsecase) Items(id int) ([]entity.ProductResponse, error) {
	entities, err := uc.repo.GetItems(id)

	if err != nil {
		return nil, err
	}
	product := []entity.ProductResponse{}

	for _, val := range entities {
		p := entity.ProductResponse{
			ID:     val.ID,
			Name:   val.Name,
			Slug:   val.Slug,
			Status: val.Status,
			Items:  []entity.ProductItemsResponse{},
		}
		// for _, item := range val.I {
		// 	p.Items = append(p.Items, entity.ProductItemsResponse{
		// 		ID:     item.ID,
		// 		Name:   item.Name,
		// 		Slug:   item.Slug,
		// 		Status: item.Status,
		// 	})
		// }
		product = append(product, p)
	}

	return product, nil
}

// func (uc *ProductUsecase) ItemInfo(product_id int, product_item_id int) (entity.ProductResponse, error) {
// 	entities, err := uc.repo.GetItemInfo(product_id, product_item_id)

// 	if err != nil {
// 		return entity.ProductResponse{}, err
// 	}
// 	product := entity.ProductResponse{
// 		ID:                   entities.ID,
// 		Name:                 entities.Name,
// 		Slug:                 entities.Slug,
// 		AdditionalInfomation: entities.AdditionalInfomation,
// 		SupportInfomation:    entities.SupportInfomation,
// 		Description:          entities.Description,
// 		Status:               entities.Status,
// 		Attributes:           []entity.ProductAttributesResponse{},
// 		Items:                []entity.ProductItemsResponse{},
// 	}
// 	for _, attribute := range entities.ProductAttributes {
// 		product.Attributes = append(product.Attributes, entity.ProductAttributesResponse{
// 			ID:     attribute.ID,
// 			Name:   attribute.Name,
// 			Key:    attribute.Key,
// 			Value:  attribute.Value,
// 			Status: attribute.Status,
// 		})
// 	}
// 	for key, item := range entities.ProductItems {
// 		product.Items = append(product.Items, entity.ProductItemsResponse{
// 			ID:         item.ID,
// 			Name:       item.Name,
// 			Slug:       item.Slug,
// 			Status:     item.Status,
// 			Attributes: []entity.ProductAttributesResponse{},
// 			Images:     []entity.ProductItemImagesResponse{},
// 			Colors:     []entity.ProductItemColorsResponse{},
// 		})

// 		for _, attribute := range item.ProductAttributes {
// 			product.Items[key].Attributes = append(product.Items[key].Attributes, entity.ProductAttributesResponse{
// 				ID:     attribute.ID,
// 				Name:   attribute.Name,
// 				Key:    attribute.Key,
// 				Value:  attribute.Value,
// 				Status: attribute.Status,
// 			})
// 		}
// 		for _, image := range item.ProductItemImages {
// 			product.Items[key].Images = append(product.Items[key].Images, entity.ProductItemImagesResponse{
// 				ID:   image.ID,
// 				Name: image.Name,
// 			})
// 		}

// 		for _, color := range item.ProductItemColors {
// 			product.Items[key].Colors = append(product.Items[key].Colors, entity.ProductItemColorsResponse{
// 				ID:       color.ID,
// 				Name:     color.Name,
// 				Hexcode:  color.Hexcode,
// 				Price:    color.Price,
// 				Discount: color.Discount,
// 				Quantity: color.Quantity,
// 				Status:   color.Status,
// 			})
// 		}
// 	}

// 	return product, nil
// }
