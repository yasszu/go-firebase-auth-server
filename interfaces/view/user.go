package view

import "go-firebase-auth-server/domain/entity"

type User struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewUser(user *entity.User) *User {
	return &User{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}
