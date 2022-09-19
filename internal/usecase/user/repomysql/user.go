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
