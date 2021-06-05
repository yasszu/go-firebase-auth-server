package repository

import "go-firebase-auth-server/domain/entity"

type AccountRepository interface {
	GetAccountByEmail(email string) (*entity.Account, error)
	GetAccountByID(accountID uint) (*entity.Account, error)
	CreateAccount(account *entity.Account) error
	UpdateAccount(account *entity.Account) error
	DeleteAccount(accountID uint) error
	RegisterFirebaseUser(user *entity.FirebaseUser) (*entity.Account, error)
}
