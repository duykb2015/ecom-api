package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

func (u *UserUsecase) GetCart(userID int) ([]entity.CartResponse, error) {
	cart, err := u.r.GetCart(userID)
	resp := make([]entity.CartResponse, 0)

	if err != nil {
		return nil, err
	}
	for _, item := range cart {
		productItem, err := u.r.GetProductItem(item.ID)
		if err != nil {
			return nil, err
		}

		productItemColors, err := u.r.GetProductItemColors(item.ID)
		if err != nil {
			return nil, err
		}
		productItemColorsResp := make([]entity.ProductItemColorsResponse, 0)
		for _, color := range productItemColors {
			productItemColorsResp = append(productItemColorsResp, entity.ProductItemColorsResponse{
				ID:       color.ID,
				Name:     color.Name,
				Hexcode:  color.Hexcode,
				Price:    color.Price,
				Discount: color.Discount,
				Quantity: color.Quantity,
				Status:   color.Status,
			})
		}

		productItemImage, err := u.r.GetProductItemImage(item.ID)
		if err != nil {
			return nil, err
		}
		productItemImageResp := make([]entity.ProductItemImagesResponse, 0)
		productItemImageResp = append(productItemImageResp, entity.ProductItemImagesResponse{
			ID:   productItemImage[0].ID,
			Name: productItemImage[0].Name,
		})

		productItemResp := entity.ProductItemsResponse{
			ID:     productItem.ID,
			Name:   productItem.Name,
			Slug:   productItem.Slug,
			Status: productItem.Status,
			Colors: productItemColorsResp,
			Images: productItemImageResp,
		}
		resp = append(resp, entity.CartResponse{
			ID:            item.ID,
			UserID:        item.UserID,
			ProductItem:   productItemResp,
			SelectedColor: item.ColorID,
			Status:        item.Status,
		})
	}

	return resp, nil
}

func (u *UserUsecase) AddItemToCart(cartInfo entity.CartRequest) bool { return true }
func (u *UserUsecase) UpdateCart(cartInfo entity.CartRequest) bool    { return true }
func (u *UserUsecase) DeleteCart(cartID int) bool                     { return true }
