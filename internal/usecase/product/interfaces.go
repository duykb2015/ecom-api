package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type (
	// ProductUsecase interface
	ProductRepo interface {
		Get(productID int) (entity.Product, error)
		GetAll() ([]entity.Product, error)
		GetByCategory(categoryID int) ([]entity.Product, error)
		GetItems() ([]entity.ProductItems, error)
		GetItemsByID(productID int) ([]entity.ProductItems, error)
		GetAttributes(productID int) ([]entity.ProductAttributes, error)
		GetItemAttributes(productItemID int) ([]entity.ProductAttributes, error)
		GetItemImages(itemId int) ([]entity.ProductItemImages, error)
		GetItemColors(itemId int) ([]entity.ProductItemColors, error)
	}

	Product interface {
		Get() ([]entity.ProductResponse, error)
		Category(categoryID int) ([]entity.ProductResponse, error)
		Items(ProductID int, ItemID int) (entity.ProductResponse, error)
		GetHotDeal() ([]entity.ProductBasicInfoResponse, error)
		GetLine() ([]entity.ProductResponse, error)
	}
)
