package entity

import "time"

type User struct {
	ID          uint      `gorm:"primary_key"`
	UserName    string    `gorm:"username"`
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

type AuthResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg,omitempty"`
	Token   string `json:"token,omitempty"`
}
type UserInfoResponse struct {
	ID       uint   `json:"primary_key"`
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    int    `json:"phone"`
	Token    string `json:"token"`
	Status   uint   `json:"column:status"`
}

type AuthRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserName string `json:"username"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}
