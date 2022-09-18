package entity

import "time"

type User struct {
	ID          uint      `gorm:"primary_key"`
	Name        string    `gorm:"username"`
	Email       string    `gorm:"email"`
	Password    string    `gorm:"password"`
	Address     string    `gorm:"address"`
	Phone       int       `gorm:"phone"`
	Token       string    `gorm:"token"`
	Status      uint      `gorm:"column:status"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	LastLoginAt time.Time `gorm:"column:last_login_at"`
}

type UserResponse struct {
	ID    uint   `gorm:"primary_key"`
	Token string `json:"token"`
}
type UserInfoResponse struct {
	ID       uint   `gorm:"primary_key"`
	Name     string `gorm:"username"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
	Address  string `gorm:"address"`
	Phone    int    `gorm:"phone"`
	Token    string `json:"token"`
	Status   uint   `gorm:"column:status"`
}
