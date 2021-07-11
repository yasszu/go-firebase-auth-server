package entity

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	UID       string
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserResponse struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (e *User) Response() *UserResponse {
	return &UserResponse{
		UserID:   e.ID,
		Username: e.Username,
		Email:    e.Email,
	}
}
