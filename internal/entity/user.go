package entity

import "time"

type User struct {
	ID          uint      `gorm:"primary_key"`
	Name        string    `gorm:"name"`
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
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    int    `json:"phone"`
	Token    string `json:"token"`
	Status   uint   `json:"column:status"`
}

type AuthRequest struct {
	ID          string `json:"id,omitempty"`
	Email       string `json:"email,omitempty" `
	Password    string `json:"password,omitempty"`
	OldPassword string `json:"old_password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
	Name        string `json:"name,omitempty"`
	Address     string `json:"address,omitempty"`
	Phone       string `json:"phone,omitempty"`
}
