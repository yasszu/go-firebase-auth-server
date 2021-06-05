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
	AccountID uint
	UID       string
	Username  string
	Email     string
}
