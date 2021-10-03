package response

import "github.com/yasszu/go-firebase-auth-server/domain/entity"

type User struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewUser(e *entity.User) *User {
	return &User{
		UserID:   e.ID,
		Username: e.Username,
		Email:    e.Email,
	}
}
