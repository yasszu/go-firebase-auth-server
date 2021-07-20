package registry

import (
	"github.com/yasszu/go-firebase-auth-server/application/usecase"
	"github.com/yasszu/go-firebase-auth-server/domain/repository"
	"github.com/yasszu/go-firebase-auth-server/domain/service"
	"github.com/yasszu/go-firebase-auth-server/infrastructure/firebase"
	"github.com/yasszu/go-firebase-auth-server/infrastructure/persistence"
	"gorm.io/gorm"
)

type Usecase interface {
	NewIndex() usecase.IndexUsecase
	NewUser() usecase.UserUsecase
}

func NewUsecase(db *gorm.DB) Usecase {
	return &usecaseImpl{
		authenticationService: firebase.NewAuthenticationService(),
		userRepository:        persistence.NewUserRepository(db),
	}
}

type usecaseImpl struct {
	db                    *gorm.DB
	authenticationService service.AuthenticationService
	userRepository        repository.UserRepository
}

func (u *usecaseImpl) NewIndex() usecase.IndexUsecase {
	return usecase.NewIndexUsecase(u.db)
}

func (u *usecaseImpl) NewUser() usecase.UserUsecase {
	return usecase.NewUserUsecase(u.userRepository, u.authenticationService)
}
