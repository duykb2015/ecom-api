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

func (uc *ProductUsecase) Get() ([]entity.ProductResponse, error) {
	resp := make([]entity.ProductResponse, 0)

	products, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		resp = append(resp, entity.ProductResponse{
			ID:     product.ID,
			Name:   product.Name,
			Slug:   product.Slug,
			Status: product.Status,
		})
	}
	return resp, nil
}

func (uc *ProductUsecase) Category(CategoryID int) ([]entity.ProductResponse, error) {
	resp := make([]entity.ProductResponse, 0)

	products, err := uc.repo.GetByCategory(CategoryID)
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		resp = append(resp, entity.ProductResponse{
			ID:     product.ID,
			Name:   product.Name,
			Slug:   product.Slug,
			Status: product.Status,
		})
	}
	return resp, nil
}

func (uc *ProductUsecase) Items(ProductID int, ItemID int) (entity.ProductResponse, error) {

	product, err := uc.repo.Get(ProductID)
	if err != nil {
		return entity.ProductResponse{}, err
	}

	attributes, err := uc.repo.GetAttributes(ProductID)
	if err != nil {
		return entity.ProductResponse{}, err
	}
	attributeResponse := []entity.ProductAttributesResponse{}
	for _, attribute := range attributes {
		attributeResponse = append(attributeResponse, entity.ProductAttributesResponse{
			ID:     attribute.ID,
			Name:   attribute.Name,
			Key:    attribute.Key,
			Value:  attribute.Value,
			Status: attribute.Status,
		})
	}

	item, err := uc.repo.GetItems(ItemID)
	if err != nil {
		return entity.ProductResponse{}, err
	}

	images, err := uc.repo.GetItemImages(int(item.ID))
	if err != nil {
		return entity.ProductResponse{}, err
	}
	imagesResponse := []entity.ProductItemImagesResponse{}

	for _, image := range images {
		imagesResponse = append(imagesResponse, entity.ProductItemImagesResponse{
			ID:   image.ID,
			Name: image.Name,
		})
	}

	colors, err := uc.repo.GetItemColors(int(item.ID))
	if err != nil {
		return entity.ProductResponse{}, err
	}
	colorsResponse := []entity.ProductItemColorsResponse{}

	for _, color := range colors {
		colorsResponse = append(colorsResponse, entity.ProductItemColorsResponse{
			ID:       color.ID,
			Name:     color.Name,
			Hexcode:  color.Hexcode,
			Price:    color.Price,
			Discount: color.Discount,
			Quantity: color.Quantity,
			Status:   color.Status,
		})
	}

	itemAttributes, err := uc.repo.GetItemAttributes(int(item.ID))
	if err != nil {
		return entity.ProductResponse{}, err
	}
	itemAttributesResponse := []entity.ProductAttributesResponse{}
	for _, itemAttribute := range itemAttributes {
		itemAttributesResponse = append(itemAttributesResponse, entity.ProductAttributesResponse{
			ID:     itemAttribute.ID,
			Name:   itemAttribute.Name,
			Key:    itemAttribute.Key,
			Value:  itemAttribute.Value,
			Status: itemAttribute.Status,
		})
	}

	itemResponse := make([]entity.ProductItemsResponse, 0)
	itemResponse = append(itemResponse, entity.ProductItemsResponse{
		ID:         item.ID,
		Name:       item.Name,
		Slug:       item.Slug,
		Status:     item.Status,
		Images:     imagesResponse,
		Colors:     colorsResponse,
		Attributes: itemAttributesResponse,
	})

	resp := entity.ProductResponse{
		ID:                    product.ID,
		Name:                  product.Name,
		Slug:                  product.Slug,
		AdditionalInformation: product.AdditionalInfomation,
		SupportInformation:    product.SupportInfomation,
		Description:           product.Description,
		Status:                product.Status,
		Attributes:            attributeResponse,
		Items:                 itemResponse,
	}

	return resp, nil
}
