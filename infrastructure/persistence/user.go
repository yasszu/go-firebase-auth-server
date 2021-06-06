package persistence

import (
	"gorm.io/gorm"

	"go-firebase-auth-server/domain/entity"
	"go-firebase-auth-server/domain/repository"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

func (r UserRepository) Crete(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r UserRepository) GetByUID(uid string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("uid = ?", uid).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
