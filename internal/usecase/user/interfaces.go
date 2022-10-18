package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type UserRepo interface {
	GetInfo(tokenString string) (entity.User, error)
	GetByID(id int) (entity.User, error)
	GetByEmail(email string) (entity.User, error)

	Create(insertData map[string]interface{}) error
	Update(userID int, updateData map[string]interface{}) error
	SaveToken(userID uint, tokenString string) error

	GetCart(userID int) ([]entity.Cart, error)
	GetProductItem(ProductItemId int) (entity.ProductItems, error)
	GetProductItemColors(ProductItemId int) ([]entity.ProductItemColors, error)
	GetProductItemImage(ProductItemId int) ([]entity.ProductItemImages, error)
}

type User interface {
	AuthLogin(email string, password string) (entity.AuthResponse, error)
	AuthRegister(request entity.AuthRequest) (entity.AuthResponse, error)
	GetInfo(tokenString string) (entity.UserInfoResponse, error)
	UpdateInfo(request entity.AuthRequest) error

	GetCart(userID int) ([]entity.CartResponse, error)
	AddItemToCart(cartInfo entity.CartRequest) bool
	UpdateCart(cartInfo entity.CartRequest) bool
	DeleteCart(cartID int) bool
}
