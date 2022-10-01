package repomysql

import (
	"github.com/duykb2015/ecom-api/internal/entity"
	"gorm.io/gorm"
)

type CartRepo struct {
	*gorm.DB
}

func NewCart(db *gorm.DB) *CartRepo {
	return &CartRepo{db}
}

func (r *CartRepo) GetByUserID(userID int) ([]entity.Cart, error) {
	cart := []entity.Cart{}
	err := r.Table("cart").Where("user_id = ?", userID).Find(&cart).Error
	return cart, err
}
