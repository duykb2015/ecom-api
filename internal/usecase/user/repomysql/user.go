package repomysql

import "gorm.io/gorm"

type UserRepo struct {
	*gorm.DB
}

func New(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}
