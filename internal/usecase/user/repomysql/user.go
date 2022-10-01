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

func (u *UserRepo) GetByID(id int) (entity.User, error) {
	user := entity.User{}
	err := u.Table("user").Where("id = ?", id).Find(&user).Error
	return user, err
}

func (u *UserRepo) GetByEmail(email string) (entity.User, error) {
	user := entity.User{}
	err := u.Table("user").Where("email = ?", email).Find(&user).Error
	return user, err
}

func (u *UserRepo) SaveToken(userID uint, tokenString string) error {
	err := u.Table("user").Where("id = ?", userID).Update("token", tokenString).Error
	return err
}

func (u *UserRepo) Create(insertData map[string]interface{}) error {
	err := u.Table("user").Create(insertData).Error
	return err
}

func (u *UserRepo) Update(userID int, updateData map[string]interface{}) error {
	err := u.Table("user").Where("id = ?", userID).Updates(updateData).Error
	return err
}

func (u *UserRepo) GetInfo(tokenString string) (entity.User, error) {
	user := entity.User{}
	err := u.Table("user").Where("token = ?", tokenString).First(&user).Error
	return user, err
}
