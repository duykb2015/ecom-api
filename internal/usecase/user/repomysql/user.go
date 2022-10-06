package repomysql

import (
	"github.com/duykb2015/ecom-api/internal/entity"
	"gorm.io/gorm"
)

const (
	userTable              = "user"
	cartTalbe              = "cart"
	productItemTable       = "product_items"
	productItemColorsTable = "product_item_colors"
	productItemImagesTable = "product_item_images"
)

type UserRepo struct {
	*gorm.DB
}

func New(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (u *UserRepo) GetByID(id int) (entity.User, error) {
	user := entity.User{}
	err := u.Table(userTable).Where("id = ?", id).Find(&user).Error
	return user, err
}

func (u *UserRepo) GetByEmail(email string) (entity.User, error) {
	user := entity.User{}
	err := u.Table(userTable).Where("email = ?", email).Find(&user).Error
	return user, err
}

func (u *UserRepo) SaveToken(userID uint, tokenString string) error {
	err := u.Table(userTable).Where("id = ?", userID).Update("token", tokenString).Error
	return err
}

func (u *UserRepo) Create(insertData map[string]interface{}) error {
	err := u.Table(userTable).Create(insertData).Error
	return err
}

func (u *UserRepo) Update(userID int, updateData map[string]interface{}) error {
	err := u.Table(userTable).Where("id = ?", userID).Updates(updateData).Error
	return err
}

func (u *UserRepo) GetInfo(tokenString string) (entity.User, error) {
	user := entity.User{}
	err := u.Table(userTable).Where("token = ?", tokenString).First(&user).Error
	return user, err
}

func (r *UserRepo) GetCart(userID int) ([]entity.Cart, error) {
	cart := []entity.Cart{}
	err := r.Table(cartTalbe).Where("user_id = ?", userID).Find(&cart).Error
	return cart, err
}

func (r *UserRepo) GetProductItem(ProductItemId int) (entity.ProductItems, error) {
	productItem := entity.ProductItems{}
	err := r.Table(productItemTable).Where("id = ?", ProductItemId).Find(&productItem).Error
	return productItem, err
}

func (r *UserRepo) GetProductItemColors(ProductItemId int) ([]entity.ProductItemColors, error) {
	productItemColors := []entity.ProductItemColors{}
	err := r.Table(productItemColorsTable).Where("product_item_id = ?", ProductItemId).Find(&productItemColors).Error
	return productItemColors, err
}

func (r *UserRepo) GetProductItemImage(ProductItemId int) ([]entity.ProductItemImages, error) {
	productItemImage := []entity.ProductItemImages{}
	err := r.Table(productItemImagesTable).Where("product_item_id = ?", ProductItemId).First(&productItemImage).Error
	return productItemImage, err
}
