package repository

import "go-firebase-auth-server/domain/entity"

//go:generate mockgen -source=./user.go -destination=./mock/user.go -package=mock
type UserRepository interface {
	Crete(user *entity.User) error
	GetByUID(uid entity.UID) (*entity.User, error)
}
