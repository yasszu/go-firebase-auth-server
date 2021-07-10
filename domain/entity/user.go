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

type userResponse struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (e *User) Response() *userResponse {
	return &userResponse{
		UserID:   e.ID,
		Username: e.Username,
		Email:    e.Email,
	}
}
