package repository

import "go-firebase-auth-server/domain/entity"

type UserRepository interface {
	Crete(user *entity.User) error
	GetByUID(uid string) (*entity.User, error)
}
