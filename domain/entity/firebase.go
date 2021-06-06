package entity

import "time"

type FirebaseAuthentication struct {
	AccountID uint
	UID       string
	CreatedAt time.Time
	UpdatedAt time.Time

	Account *Account
}

type FirebaseUser struct {
	UID      string
	Username string
	Email    string
}

func (u *FirebaseUser) User() *User {
	return &User{
		UID:      u.UID,
		Username: u.Username,
		Email:    u.Email,
	}
}
