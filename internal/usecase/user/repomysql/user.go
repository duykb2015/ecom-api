package repomysql

import (
	"github.com/duykb2015/ecom-api/internal/entity"
	"gorm.io/gorm"
)

type UserRepo struct {
	*gorm.DB
}

func New(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (u *UserRepo) GetUser(email string) (entity.User, error) {
	user := entity.User{}
	err := u.Table("user").Where("email = ?", email).Find(&user).Error
	return user, err
}

func (u *UserRepo) SaveToken(userID uint, tokenString string) error {
	err := u.Table("user").Where("id = ?", userID).Update("token", tokenString).Error
	return err
}

func (u *UserRepo) CreateUser(insertData map[string]interface{}) error {
	err := u.Table("user").Create(insertData).Error
	return err
}
